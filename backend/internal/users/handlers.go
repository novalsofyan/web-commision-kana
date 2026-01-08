// Package users menyediakan handler untuk autentikasi
package users

import (
	auth "backend-web-commision-kana/internal/middleware/auth"
	"backend-web-commision-kana/internal/utils/jsonresp"
	"encoding/json"
	"log/slog"
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

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
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

	// Set HTTP-only cookie dengan expiry 30 hari
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    res.Token,
		Path:     "/",
		MaxAge:   30 * 24 * 60 * 60,
		HttpOnly: true,
		Secure:   false, // Set true di production dengan HTTPS
		SameSite: http.SameSiteLaxMode,
	})

	// Response tanpa token (sudah di cookie)
	h.resp.WriteData(w, http.StatusOK, map[string]string{
		"username": res.Username,
		"message":  "login berhasil",
	})
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	token, ok := auth.GetToken(r)
	if !ok {
		h.resp.WriteData(w, http.StatusUnauthorized, map[string]string{
			"error": "Unauthorized",
		})
		return
	}

	res, err := h.svc.Logout(r.Context(), token)
	if err != nil {
		h.resp.WriteData(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	// Clear cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	h.resp.WriteData(w, http.StatusOK, res)
}

func (h *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var req ReqUpdateProfile
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		h.resp.WriteData(w, http.StatusBadRequest, map[string]string{"error": "Bad request"})
		return
	}
	userID, ok := r.Context().Value(auth.UserIDKey).(int32)
	if !ok {
		h.resp.WriteData(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		return
	}
	res, err := h.svc.UpdateProfile(r.Context(), userID, req)
	if err != nil {
		slog.Error("Error update profile", "err", err)
		h.resp.WriteData(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	h.resp.WriteData(w, http.StatusOK, res)
}
