package auth

import "net/http"

func GetUserID(r *http.Request) (int64, bool) {
	userID, ok := r.Context().Value(UserIDKey).(int64)
	return userID, ok
}

func GetToken(r *http.Request) (string, bool) {
	token, ok := r.Context().Value(TokenKey).(string)
	return token, ok
}
