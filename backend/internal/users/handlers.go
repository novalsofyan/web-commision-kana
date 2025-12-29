package users

import (
	"backend-web-commision-kana/internal/utils/jsonresp"
	"encoding/json"
	"net/http"
)

type Handler struct {
	svc  Service
	resp jsonresp.JSONResponder
}

func NewHandler(svc Service, resp jsonresp.JSONResponder) *Handler {
	return &Handler{svc: svc, resp: resp}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin

	// 1. Decode JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.resp.WriteData(w, http.StatusBadRequest, map[string]string{
			"error": "request tidak valid",
		})
		return
	}

	res, err := h.svc.Login(r.Context(), req)
	if err != nil {
		h.resp.WriteData(w, http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
		return
	}

	h.resp.WriteData(w, http.StatusOK, res)
}
