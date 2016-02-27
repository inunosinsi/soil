package session

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFlashSession(t *testing.T) {

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "", nil)

	token1, _ := GetFlashSession(w, r)
	token2, go_token := GetFlashSession(w, r)

	log.Println(token1)
	log.Println(token2)
	log.Println(go_token)

	if token1 != go_token {
		t.Error("フラッシュセッションの値が一致しません")
	}

	if token1 == token2 {
		t.Error("一回目のトークンの生成と二回目の値が一致しています")
	}
}
