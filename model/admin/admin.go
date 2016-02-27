package admin

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"../../dbconf"
)

type Admin struct {
	id       int
	login_id string
	password string
}

func Check() bool {

	conf := dbconf.GetDBConfig()

	db, err := sql.Open("mysql", conf.User+":"+conf.Pass+"@/"+conf.Db)
	if err != nil {
		db.Close()
		return false
	}

	res, err := db.Query("SELECT id FROM Administrator LIMIT 1")
	if err != nil {
		db.Close()
		return false
	}

	for res.Next() {
		var id int
		err = res.Scan(&id)
		if err != nil {
			db.Close()
			return false
		}

		//データがあればtrueを返す
		if id > 0 {
			break
		}
	}

	return true
}

func Insert(loginId interface{}, password interface{}) int64 {

	conf := dbconf.GetDBConfig()

	db, err := sql.Open("mysql", conf.User+":"+conf.Pass+"@/"+conf.Db)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	//データベースに値を突っ込んでみる
	stmt, err := db.Prepare("INSERT Administrator SET login_id=?,password=?")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(loginId, password)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return id
}

func GetPasswordHashByLoginId(loginId string) string {
	var passHash string

	conf := dbconf.GetDBConfig()

	db, err := sql.Open("mysql", conf.User+":"+conf.Pass+"@/"+conf.Db)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	stmt, err := db.Prepare("SELECT password FROM Administrator WHERE login_id = ?")
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query(loginId)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&passHash)
		if err != nil {
			panic(err)
		}
	}

	return passHash
}
