package main

import (
	"backend-web-commision-kana/internal/utils/jsonresp"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Addr    string
	Logger  *slog.Logger
	JSONres jsonresp.JSONResponder
	DbConf  *DBConfig
}

type DBConfig struct {
	DSN      string
	maxConns int
	minConns int
}

type Application struct {
	Conf     *Config
	DBConfig *pgxpool.Pool
}
