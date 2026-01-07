// Package users menyediakan type untuk autentikasi
package users

type ReqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResLogin struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

type ReqLogout struct {
	Token string `json:"token"`
}

type ResLogout struct {
	Msg string `json:"msg"`
}

type ReqUpdateProfile struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type ResUpdateProfile struct {
	Username string `json:"username"`
}
