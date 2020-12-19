package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mosk-service/Dao"
	"mosk-service/Database"
	"mosk-service/Loader"
	"mosk-service/Router"
	"strconv"
)

func main() {
	// 加载配置文件
	conf := Loader.LoadConfig()

	mysql := Database.MysqlConnection(
		conf.Mysql.Server,
		conf.Mysql.Port,
		conf.Mysql.User,
		conf.Mysql.Passwd,
		conf.Mysql.Database,
		conf.Mysql.MaxOpenConn,
		conf.Mysql.MaxIdleConn)

	Dao.GetConfigBySetting("base")
	fmt.Println(mysql,Database.MysqlDB)

	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	// 注册全局路由
	Router.SetupRouter(r)

	r.GET("/ncda/basicinfo/region/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "OK",
			"data": gin.H{
				"id":                1,
				"cityName":          "北京市",
				"divisionId":        19,
				"divisionName":      "测试",
				"divisionLongitude": "116.403039",
				"divisionLatitude":  "39.914466",
				"name":              "测试",
				"regionCode":        "test",
				"longitude":         "116.403039",
				"latitude":          "39.914466",
				"textLongitude":     "116.403039",
				"textLatitude":      "39.914466",
				"deptId":            15,
			},
		})
	})

	_ = r.Run(":" + strconv.Itoa(conf.Port))
}
