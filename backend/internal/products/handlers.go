// Package products for showing products (search products etc.)
package products

import (
	"backend-web-commision-kana/internal/middleware/auth"
	"backend-web-commision-kana/internal/utils/jsonresp"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
)

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

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req ReqAdminProducts

	userID, ok := auth.GetUserID(r)
	if !ok {
		h.resp.WriteData(w, http.StatusUnauthorized, map[string]string{
			"error": "Unauthorized",
		})
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		h.resp.WriteData(w, http.StatusBadRequest, map[string]string{
			"error": "Bad Request",
		})
		return
	}

	res, err := h.svc.CreateProduct(r.Context(), req, userID)
	if err != nil {
		slog.Error("Internal server error", "error", err)
		h.resp.WriteData(w, http.StatusInternalServerError, map[string]string{
			"error": "Internal server error",
		})
		return
	}

	h.resp.WriteData(w, http.StatusOK, res)
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("product_id")

	productID, err := strconv.Atoi(idStr)
	if err != nil {
		h.resp.WriteData(w, http.StatusBadRequest, map[string]string{
			"error": "Bad Request",
		})
		return
	}

	res, err := h.svc.DeleteProduct(r.Context(), int32(productID))
	if err != nil {
		if errors.Is(err, ErrProductNotFound) {
			h.resp.WriteData(w, http.StatusNotFound, map[string]string{
				"error": "Product not found",
			})
			return
		}

		slog.Error("Internal server error", "error", err)
		h.resp.WriteData(w, http.StatusInternalServerError, map[string]string{
			"error": "Internal server error",
		})
		return
	}

	h.resp.WriteData(w, http.StatusOK, res)
}

func (h *Handler) GetProductAdmin(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserID(r)
	if !ok {
		h.resp.WriteData(w, http.StatusUnauthorized, map[string]string{
			"error": "Unauthorized",
		})
		return
	}

	res, err := h.svc.GetProductAdmin(r.Context(), userID)
	if err != nil {
		slog.Error("Internal server error", "error", err)
		h.resp.WriteData(w, http.StatusInternalServerError, map[string]string{
			"error": "Internal server error",
		})
		return
	}

	h.resp.WriteData(w, http.StatusOK, res)
}
