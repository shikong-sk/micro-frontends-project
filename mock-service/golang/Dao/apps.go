package Dao

import (
	"mock-service/Database"
)

type App struct {
	Name       string `db:"name" json:"name"`
	Entry      string `db:"entry" json:"entry"`
	Container  string `db:"container" json:"container"`
	ActiveRule string `db:"activeRule" json:"activeRule"`
	Enable     bool   `db:"enable" json:"enable"`
}

func GetApps() []App {
	var data []App
	sql := "SELECT name, entry, container, activeRule, enable FROM apps"
	_ = Database.MysqlDB.Select(&data, sql)
	return data
}

func GetAppsByName(name string) []App {
	var data []App
	sql := "SELECT name, entry, container, activeRule, enable FROM apps where name = ?"
	_ = Database.MysqlDB.Select(&data, sql, name)
	return data
}
