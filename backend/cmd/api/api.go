package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v3"
)

func (api *Application) mount() http.Handler {
	r := chi.NewRouter()

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

	// userHandler := users.NewHandler()

	// r.Route("/api/", func(r chi.Router) {
	// 	r.Route("/users", func(r chi.Router) {
	// 		r.Post("/login", )
	// 	}
	// })

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
