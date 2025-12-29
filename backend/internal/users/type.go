package users

import (
	"backend-web-commision-kana/internal/utils/jsonresp"
	"context"
)

type Service interface {
	Login(ctx context.Context, req ReqLogin) (*ResLogin, error)
}

type ReqLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ResLogin struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

type Res struct {
	JSONres jsonresp.JSONResponder
}
