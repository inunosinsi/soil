package goydb

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Goydb interface {
	TableName() string
}

func Insert(s Goydb) (int64, error) {
	//構造体であるか調べる
	if v := reflect.Indirect(reflect.ValueOf(s)); v.Kind().String() != "struct" {
		return 0, errors.New("Not Struct")
	}

	m := makeMap(s)

	db := Conn()
	defer db.Close()

	tbName := s.TableName()
	q := "INSERT " + tbName + " SET "
	c := 0

	values := make([]interface{}, 0)
	for key, v := range m {
		if key == "id" || key == "Id" {
			continue
		}
		
		if c > 0 {
			q += ", "
		}
		if strings.Index(key, "D") > 0 {
			key = strings.Replace(key, "D", "_d", 1)
		} else if strings.Index(key, "I") > 0 {
			key = strings.Replace(key, "I", "_i", 1)
		} else if strings.Index(key, "K") > 0 {
			key = strings.Replace(key, "K", "_k", 1)
		}
		q += strings.ToLower(key) + "=?"
		values = append(values, v)
		c++
	}

	//データベースに値を突っ込んでみる
	stmt, err := db.Prepare(q)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(values...)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func Update(s Goydb) error {
	//構造体であるか調べる
	if v := reflect.Indirect(reflect.ValueOf(s)); v.Kind().String() != "struct" {
		return errors.New("Not Struct")
	}

	m := makeMap(s)

	db := Conn()
	defer db.Close()

	tbName := s.TableName()
	q := "UPDATE " + tbName + " SET "
	c := 0

	values := make([]interface{}, 0)
	for key, v := range m {
		
		if c > 0 {
			q += ", "
		}
		if strings.Index(key, "D") > 0 {
			key = strings.Replace(key, "D", "_d", 1)
		} else if strings.Index(key, "I") > 0 {
			key = strings.Replace(key, "I", "_i", 1)
		} else if strings.Index(key, "K") > 0 {
			key = strings.Replace(key, "K", "_k", 1)
		}
		q += strings.ToLower(key) + "=?"
		values = append(values, v)
		c++
	}

	//データベースに値を突っ込んでみる
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(values...)
	return err
}

func Conn() *sql.DB {
	conf := getConfig()

	db, err := sql.Open("mysql", conf.User+":"+conf.Pass+"@/"+conf.Db)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func makeMap(s interface{}) map[string]interface{} {
	m := make(map[string]interface{})

	v := reflect.Indirect(reflect.ValueOf(s))
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)

		if field.CanSet() {
			m[t.Field(i).Name] = field.Interface()
		}
	}

	return m
}

type Config struct {
	Db   string
	User string
	Pass string
}

func getConfig() Config {
	var c Config

	p, _ := os.Getwd()
	if strings.Index(p, "\\") > 0 {
		p = strings.Replace(p, "\\", "/", -1)
	}
	p = strings.Replace(p, "/dbconf", "", 1)

	jsonString, err := ioutil.ReadFile(p + "/conf/mysql.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonString, &c)
	if err != nil {
		panic(err)
	}

	return c
}
