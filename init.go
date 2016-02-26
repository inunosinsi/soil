package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"net/http"
	
	"./dbconf"
)

type initHandler struct {
	next http.Handler
}

func (h *initHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var doRedirect = false

	conf := dbconf.GetDBConfig()

	db, err := sql.Open("mysql", conf.User+":"+conf.Pass+"@/"+conf.Db)
	if err != nil {
		doRedirect = true
	}

	_, err = db.Query("SELECT id FROM administrator LIMIT 1")
	if err != nil {
		doRedirect = true
	}

	db.Close()

	//初期化フラグがtrueの場合は初期化ページへ
	if doRedirect {
		//初期化ページへ飛ぶ
		w.Header().Set("Location", "/init")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}

	// 成功。ラップされたハンドラを呼び出します
	h.next.ServeHTTP(w, r)
}
func CheckDB(handler http.Handler) http.Handler {
	return &initHandler{next: handler}
}
