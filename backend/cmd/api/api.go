package main

import (
	auth "backend-web-commision-kana/internal/middleware"
	"backend-web-commision-kana/internal/repo"
	"backend-web-commision-kana/internal/users"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v3"
	"github.com/rs/cors"
)

func (api *Application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://127.0.0.1:5173", "http://localhost:4173", "http://127.0.0.1:4173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	// Use httplog
	r.Use(httplog.RequestLogger(api.Conf.Logger, &httplog.Options{
		Level: slog.LevelInfo,
	}))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		api.Conf.JSONres.WriteData(w, http.StatusOK, map[string]any{
			"msg": "Hello there!",
		})
	})

	queries := repo.New(api.DB)
	userSvc := users.NewServices(queries, api.DB)
	userHandler := users.NewHandler(userSvc, api.Conf.JSONres)

	r.Route("/api", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", userHandler.Login)
			r.Group(func(r chi.Router) {
				r.Use(auth.AuthMiddleware(queries))
				r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
					api.Conf.JSONres.WriteData(w, http.StatusOK, map[string]any{
						"status": "authenticated",
					})
				})
				r.Post("/logout", userHandler.Logout)
			})
		})
	})

	return r
}

func (api *Application) run(h http.Handler) error {
	server := &http.Server{
		Addr:    api.Conf.Addr,
		Handler: h,
		// TimeOut Configuration
		WriteTimeout: time.Second * 20,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 30,
	}

	slog.Info("Starting server . . .", "addr", api.Conf.Addr)

	return server.ListenAndServe()
}
