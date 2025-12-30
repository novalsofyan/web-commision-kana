package main

import (
	"backend-web-commision-kana/internal/utils/jsonresp"
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/httplog/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
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

	// DB Config
	dbCfg := &DBConfig{
		DSN:      os.Getenv("GOOSE_DBSTRING"),
		maxConns: 20,
		minConns: 5,
	}

	// Web server config & listen
	cfg := &Config{
		Addr:    ":8080",
		JSONres: jsRes,
		Logger:  logger,
		DbConf:  dbCfg,
	}

	web := &Application{
		Conf: cfg,
	}

	pCfg, err := pgxpool.ParseConfig(dbCfg.DSN)
	if err != nil {
		log.Fatal("DSN error")
	}

	pCfg.MaxConns = int32(dbCfg.maxConns)
	pCfg.MinConns = int32(dbCfg.minConns)

	// buka pool pake
	pool, err := pgxpool.NewWithConfig(ctx, pCfg)
	if err != nil {
		log.Fatal("Can't connect database", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Gagal konek! Cek username/password/host: ", err)
	}

	defer pool.Close()

	web.DB = pool

	if err := web.run(web.mount()); err != nil {
		slog.Error("Server can't start!", "status", http.StatusInternalServerError, "error", err)
		os.Exit(1)
	}
}
