package Database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
)

func MysqlConnection(server string, port int, user string, passwd string, database string, maxOpenConn int, maxIdleConn int) *sqlx.DB {
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

	// 最大连接数
	db.SetMaxOpenConns(maxOpenConn)
	// 最大空闲连接数
	db.SetMaxIdleConns(maxIdleConn)

	var existDatabase bool
	_ = db.Get(&existDatabase, "SELECT COUNT(SCHEMA_NAME) > 0 FROM information_schema.SCHEMATA where SCHEMA_NAME = ?", database)

	fmt.Println("是否存在", database, "数据库", existDatabase)

	if !existDatabase {
		panic(database + " 数据库不存在")
	}

	_, _ = db.Query("use " + database)

	defer func() {
		_ = db.Close()
	}()
	return db
}
