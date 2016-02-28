package view

import (
	"net/http"

	"../login"
)

type logoutHandler struct {
	next http.Handler
}

func (h *logoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//ログインしていた場合はログアウト
	if isLoggedIn := login.IsLoggedIn(r); isLoggedIn {
		login.Logout(w, r)
	}

	w.Header().Set("Location", "/login")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func Logout(handler http.Handler) http.Handler {
	return &logoutHandler{next: handler}
}
