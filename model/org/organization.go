package org

import (
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mholt/binding"

	"../../goy/goydb"
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

func (o *Org) TableName() string {
	return "Organization"
}

func Insert(org *Org) int64 {
	var dbs goydb.Goydb = org
	id, err := goydb.Insert(dbs)
	if err != nil {
		panic(err)
	}

	return id
}

func Get(limit int) *[]Org {
	db := goydb.Conn()
	defer db.Close()

	lim := strconv.Itoa(limit)

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

	db := goydb.Conn()
	defer db.Close()

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
