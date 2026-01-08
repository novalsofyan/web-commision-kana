package main

import (
	auth "backend-web-commision-kana/internal/middleware/auth"
	"backend-web-commision-kana/internal/products"
	"backend-web-commision-kana/internal/repo"
	"backend-web-commision-kana/internal/users"
	"backend-web-commision-kana/internal/utils/scheduler"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v3"
	"github.com/go-chi/httprate"
	"github.com/rs/cors"
)

func (api *Application) mount() http.Handler {
	r := chi.NewRouter()

	rawAllowdOrigin := os.Getenv("ALLOWED_ORIGINS")
	allowedOrigins := strings.Split(rawAllowdOrigin, ",")

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(httprate.Limit(
		100,
		1*time.Minute,
		httprate.WithKeyFuncs(httprate.KeyByIP),
		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
			api.Conf.JSONres.WriteData(w, http.StatusTooManyRequests, map[string]string{
				"error": "Terlalu banyak request",
			})
		}),
	))
	r.Use(httplog.RequestLogger(api.Conf.Logger, &httplog.Options{
		Level: slog.LevelInfo,
	}))
	r.Use(middleware.Timeout(30 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		api.Conf.JSONres.WriteData(w, http.StatusOK, map[string]any{
			"msg": "Hello there!",
		})
	})

	queries := repo.New(api.DB)
	userSvc := users.NewServices(queries, api.DB)
	userHandler := users.NewHandler(userSvc, api.Conf.JSONres)

	productSvc := products.NewService(queries, api.DB)
	productHandler := products.NewHandler(productSvc, api.Conf.JSONres)

	// Initialize session cleanup scheduler (interval: 24 jam)
	api.Scheduler = scheduler.NewSessionCleanupScheduler(queries, 24*time.Hour)

	r.Route("/api", func(r chi.Router) {

		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", userHandler.Login)
			r.Group(func(r chi.Router) {
				r.Use(auth.AuthMiddleware(queries, api.Conf.JSONres))
				r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
					api.Conf.JSONres.WriteData(w, http.StatusOK, map[string]any{
						"status": "authenticated",
					})
				})
				r.Post("/logout", userHandler.Logout)
			})
		})

		r.Route("/admin", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Use(auth.AuthMiddleware(queries, api.Conf.JSONres))
				r.Get("/products", productHandler.GetProductAdmin)
				r.Post("/products", productHandler.CreateProduct)
				r.Delete("/products/{product_id}", productHandler.DeleteProduct)
				r.Patch("/me", userHandler.UpdateProfile)
			})
		})
	})

	return r
}

func (api *Application) run(h http.Handler) *http.Server {
	server := &http.Server{
		Addr:         api.Conf.Addr,
		Handler:      h,
		WriteTimeout: 20 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	slog.Info("Starting server . . .", "addr", api.Conf.Addr)

	return server
}
