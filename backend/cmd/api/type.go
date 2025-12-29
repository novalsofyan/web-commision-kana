package main

import (
	"backend-web-commision-kana/internal/utils/jsonresp"
	"log/slog"
)

type Config struct {
	Addr    string
	Logger  *slog.Logger
	JSONres jsonresp.JSONResponder
}

type Application struct {
	Conf *Config
}
