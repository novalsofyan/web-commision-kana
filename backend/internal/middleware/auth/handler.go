package auth

import "net/http"

func GetUserID(r *http.Request) (int32, bool) {
	userID, ok := r.Context().Value(UserIDKey).(int32)
	return userID, ok
}

func GetToken(r *http.Request) (string, bool) {
	token, ok := r.Context().Value(TokenKey).(string)
	return token, ok
}
