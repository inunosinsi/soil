package org

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"../../dbconf"
)

type Org struct {
	id   int
	name string
}

func Insert(name interface{}) int64 {

	conf := dbconf.GetDBConfig()

	db, err := sql.Open("mysql", conf.User+":"+conf.Pass+"@/"+conf.Db)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	//データベースに値を突っ込んでみる
	stmt, err := db.Prepare("INSERT Organization SET name=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(name)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return id
}
