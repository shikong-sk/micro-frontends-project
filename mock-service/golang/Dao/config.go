package Dao

import (
	"encoding/json"
	"fmt"
	"mock-service/Database"
)

type Config struct {
	Setting string `db:"setting"`
	Value   string `db:"value"`
}

func GetConfigBySetting(setting string) []Config {
	var data []Config
	sql := "SELECT setting, value FROM `micro-frontends`.config WHERE setting = ?"
	_ = Database.MysqlDB.Select(&data, sql, setting)
	j, _ := json.Marshal(data)
	fmt.Println(string(j))
	return data
}
