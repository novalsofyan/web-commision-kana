package main

import (
	"backend-web-commision-kana/internal/utils/jsonresp"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/httplog/v3"
	"github.com/joho/godotenv"
)

func main() {
	// Logger httplog + Elastic Common Schema
	isConcise := true
	logFormat := httplog.SchemaECS.Concise(isConcise)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: logFormat.ReplaceAttr,
	}))
	slog.SetDefault(logger)
	// JSON Response Writer
	jsRes := jsonresp.NewResponder()

	// Dotenv
	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found!", "detail", err)
	}
	env := os.Getenv("PROJECT_ENV")

	if env == "production" || env == "" {
		isConcise = false
	}

	// Web server config & listen
	cfg := &Config{
		Addr:    ":8080",
		JSONres: jsRes,
		Logger:  logger,
	}

	web := &Application{
		Conf: cfg,
	}

	if err := web.run(web.mount()); err != nil {
		slog.Error("Server can't start!", "status", http.StatusInternalServerError, "error", err)
		os.Exit(1)
	}
}
