package auth

type contextKey string

const (
	UserIDKey contextKey = "user_id"
	TokenKey  contextKey = "token"
)
