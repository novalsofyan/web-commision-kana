// Package products for showing products (search products etc.)
package products

import "backend-web-commision-kana/internal/utils/jsonresp"

type Handler struct {
	svc  Service
	resp jsonresp.JSONResponder
}

func NewHandler(svc Service, resp jsonresp.JSONResponder) *Handler {
	return &Handler{
		svc:  svc,
		resp: resp,
	}
}
