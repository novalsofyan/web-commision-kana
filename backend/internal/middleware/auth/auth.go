// Package auth for middleware auth
package auth

import (
	"backend-web-commision-kana/internal/repo"
	"backend-web-commision-kana/internal/utils/jsonresp"
	"context"
	"net/http"
	"strings"
)

func AuthMiddleware(queries *repo.Queries, resp jsonresp.JSONResponder) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			token := strings.TrimPrefix(authHeader, "Bearer ")
			token = strings.TrimSpace(token)

			if token == "" {
				resp.WriteData(w, http.StatusUnauthorized, map[string]string{
					"error": "Unauthorized",
				})
				return
			}

			userID, err := queries.SelectUserBySession(r.Context(), token)
			if err != nil {
				resp.WriteData(w, http.StatusUnauthorized, map[string]string{
					"error": "Unauthorized",
				})
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, UserIDKey, userID)
			ctx = context.WithValue(ctx, TokenKey, token)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
