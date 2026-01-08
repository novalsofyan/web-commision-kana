// Package auth for middleware auth
package auth

import (
	"backend-web-commision-kana/internal/repo"
	"backend-web-commision-kana/internal/utils/jsonresp"
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func AuthMiddleware(queries *repo.Queries, resp jsonresp.JSONResponder) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Read token dari cookie
			cookie, err := r.Cookie("auth_token")
			if err != nil {
				resp.WriteData(w, http.StatusUnauthorized, map[string]string{
					"error": "Unauthorized",
				})
				return
			}

			token := cookie.Value
			if token == "" {
				resp.WriteData(w, http.StatusUnauthorized, map[string]string{
					"error": "Unauthorized",
				})
				return
			}

			userID, err := queries.SelectUserBySession(r.Context(), token)
			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					resp.WriteData(w, http.StatusUnauthorized, map[string]string{
						"error": "Unauthorized",
					})
				} else {
					slog.Error("DATABASE CRITICAL ERROR", "err", err)
					resp.WriteData(w, http.StatusInternalServerError, map[string]string{
						"error": "Server error: Database error",
					})
				}
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, UserIDKey, userID)
			ctx = context.WithValue(ctx, TokenKey, token)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
