package Dao

import (
	"mock-service/Database"
)

type Config struct {
	Setting string `db:"setting" json:"setting"`
	Value   string `db:"value" json:"value"`
}

func GetConfigBySetting(setting string) []Config {
	var data []Config
	sql := "SELECT setting, value FROM config WHERE setting = ?"
	_ = Database.MysqlDB.Select(&data, sql, setting)
	return data
}
