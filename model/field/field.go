package field

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"../../dbconf"
)

type Field struct {
	id     int
	name   string
	org_id int
}

func Insert(name interface{}, org_id interface{}) int64 {

	conf := dbconf.GetDBConfig()

	db, err := sql.Open("mysql", conf.User+":"+conf.Pass+"@/"+conf.Db)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	//データベースに値を突っ込んでみる
	stmt, err := db.Prepare("INSERT Field SET name=?, org_id=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(name, org_id)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return id
}
