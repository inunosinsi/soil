package dbconf

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Db   string
	User string
	Pass string
}

func GetDBConfig() Config {
	var conf Config

	jsonString, err := ioutil.ReadFile("./conf/mysql.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonString, &conf)
	if err != nil {
		panic(err)
	}

	return conf
}
