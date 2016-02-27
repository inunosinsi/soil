package main

import (
	"net/http"

	"./login"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if isLoggedIn := login.IsLoggedIn(r); !isLoggedIn {
		// 未認証
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		// 成功。ラップされたハンドラを呼び出します
		h.next.ServeHTTP(w, r)
	}
}
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}
