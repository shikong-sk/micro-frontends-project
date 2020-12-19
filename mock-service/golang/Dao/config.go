package Dao

import (
	"fmt"
	"mosk-service/Database"
)

type Config struct {
	Setting string `db:"setting"`
	Value   string `db:"value"`
}

func GetConfigBySetting(setting string) []Config {
	fmt.Println(Database.MysqlDB.Ping() != nil)
	var data []Config
	_ = Database.MysqlDB.Select(&data, "SELECT setting, value FROM `micro-frontends`.config WHERE setting = ?", setting)
	fmt.Println(data)
	return data
}
