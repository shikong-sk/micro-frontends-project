package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mock-service/Database"
	"mock-service/Loader"
	"mock-service/Router"
	"mock-service/Session"
	"strconv"
)

func main() {
	// 加载配置文件
	conf := Loader.LoadConfig()

	fmt.Println(Session.Store)

	func() {
		defer func() {
			err := recover()
			if err != nil {
				_, _ = fmt.Fprintln(gin.DefaultErrorWriter, err)
			}
		}()

		Database.MysqlConnection(
			conf.Mysql.Server,
			conf.Mysql.Port,
			conf.Mysql.User,
			conf.Mysql.Passwd,
			conf.Mysql.Database,
			conf.Mysql.MaxOpenConn,
			conf.Mysql.MaxIdleConn)
	}()

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	Session.InitSession(r, conf.Session.Secret)

	// 注册全局路由
	Router.SetupRouter(r)

	_ = r.Run(":" + strconv.Itoa(conf.Port))
}
