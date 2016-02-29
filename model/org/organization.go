package org

import (
	"database/sql"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mholt/binding"

	"../../dbconf"
)

type Org struct {
	Id   int
	Name string
}

func NewOrg() Org {
	return Org{}
}

func (o *Org) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&o.Id:   "id",
		&o.Name: "name",
	}
}

func Insert(org *Org) int64 {

	conf := dbconf.GetDBConfig()

	db, err := sql.Open("mysql", conf.User+":"+conf.Pass+"@/"+conf.Db)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	o := reflect.Indirect(reflect.ValueOf(org))
	t := o.Type()

	m := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		field := o.Field(i)

		if field.CanSet() {
			m[t.Field(i).Name] = field.Interface()
		}
	}

	sql := "INSERT Organization SET "
	c := 0
	values := make([]interface{}, 0)
	for key, v := range m {
		if key == "id" || key == "Id" {
			continue
		}
		if c > 0 {
			sql += ", "
			c++
		}
		sql += strings.ToLower(key) + "=?"
		values = append(values, v)
	}

	//データベースに値を突っ込んでみる
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(values...)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return id

}

func Get(limit int) *[]Org {
	conf := dbconf.GetDBConfig()

	lim := strconv.Itoa(limit)

	db, err := sql.Open("mysql", conf.User+":"+conf.Pass+"@/"+conf.Db)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("SELECT * FROM Organization LIMIT " + lim)
	if err != nil {
		panic(err.Error())
	}

	//SQLで結果の取得数を調べてから配列を用意
	list := make([]Org, 0)

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		if id > 0 {
			list = append(list, Org{id, name})
		}
	}

	return &list
}

func GetById(orgId int) *Org {
	var org Org

	conf := dbconf.GetDBConfig()

	db, err := sql.Open("mysql", conf.User+":"+conf.Pass+"@/"+conf.Db)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	stmt, err := db.Prepare("SELECT * FROM Organization WHERE id = ?")
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query(orgId)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		org = Org{id, name}
	}

	return &org
}
