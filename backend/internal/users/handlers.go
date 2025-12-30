package users

import (
	"backend-web-commision-kana/internal/utils/jsonresp"
	"encoding/json"
	"net/http"
	"strings"
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

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	// 1. Ambil token dari Header (karena dikirim via Vue/Axios di header)
	authHeader := r.Header.Get("Authorization")

	// Kita bersihkan prefix "Bearer " (jika ada)
	token := strings.TrimPrefix(authHeader, "Bearer ")
	token = strings.TrimSpace(token)

	// 2. Panggil Service Logout
	// Kita bungkus token ke struct ReqLogout sesuai kontrak service kamu
	res, err := h.svc.Logout(r.Context(), ReqLogout{Token: token})
	if err != nil {
		h.resp.WriteData(w, http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
		return
	}

	// 3. Response Berhasil
	h.resp.WriteData(w, http.StatusOK, res)
}
