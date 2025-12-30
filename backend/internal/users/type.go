// Package users menyediakan type untuk autentikasi
package users

import (
	"backend-web-commision-kana/internal/utils/jsonresp"
	"context"
)

type Service interface {
	Login(ctx context.Context, req ReqLogin) (*ResLogin, error)
	Logout(ctx context.Context, req ReqLogout) (*ResLogout, error)
}

type ReqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResLogin struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

type ReqLogout struct {
	Token string `json:"token"`
}

type ResLogout struct {
	Msg string `json:"msg"`
}

type Res struct {
	JSONres jsonresp.JSONResponder
}
