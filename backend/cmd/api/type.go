package main

import (
	"backend-web-commision-kana/internal/utils/jsonresp"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBConfig struct {
	DSN      string
	maxConns int
	minConns int
}

type Config struct {
	Addr    string
	Logger  *slog.Logger
	JSONres jsonresp.JSONResponder
	DBConf  *DBConfig
}

type Application struct {
	Conf *Config
	DB   *pgxpool.Pool
}
