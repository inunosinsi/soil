package field

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"../../dbconf"
)

type Field struct {
	Id    int
	Name  string
	OrgId int
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

func Get(limit int) *[]Field {
	conf := dbconf.GetDBConfig()

	lim := strconv.Itoa(limit)

	db, err := sql.Open("mysql", conf.User+":"+conf.Pass+"@/"+conf.Db)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("SELECT * FROM Field LIMIT " + lim)
	if err != nil {
		panic(err.Error())
	}

	//SQLで結果の取得数を調べてから配列を用意
	list := make([]Field, 0)

	for rows.Next() {
		var id int
		var name string
		var orgId int
		err = rows.Scan(&id, &name, &orgId)
		if err != nil {
			panic(err)
		}
		if id > 0 {
			list = append(list, Field{id, name, orgId})
		}
	}

	return &list
}
