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

func NewHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.resp.WriteData(w, http.StatusBadRequest, map[string]string{
			"error": "request tidak valid",
		})
	}
}
