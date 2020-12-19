package Database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
)

func MysqlConnection(server string, port int, user string, passwd string, database string) *sqlx.DB {
	dataSource := user + ":" + passwd + "@tcp(" + server + ":" + strconv.Itoa(port) + ")/"
	fmt.Println("dataSource", dataSource)
	db, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	} else {
		err = db.Ping()
		if err != nil {
			panic(err)
		}
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)

	var existDatabase bool
	_ = db.Select(existDatabase, "SELECT COUNT(SCHEMA_NAME) FROM information_schema.SCHEMATA where SCHEMA_NAME = ?", database)

	fmt.Println("是否存在", database, "数据库", existDatabase)
	defer func() {
		_ = db.Close()
	}()
	return db
}
