package Controller

import (
	"github.com/gin-gonic/gin"
	"mock-service/Common/Error"
	"mock-service/Dao"
	"mock-service/Database"
)

func SetupBaseRouter(r *gin.Engine) {
	router := r.Group("/base")
	getConfig(router)
	getApps(router)
	getAppsWithName(router)
}

func getConfig(r *gin.RouterGroup) {
	r.GET("/getConfig", func(context *gin.Context) {
		if Database.MysqlDB != nil {
			data := Dao.GetConfigBySetting("base")
			context.JSON(200, data)
		} else {
			context.JSON(500, Error.DatabaseConnectError)
		}
	})
}

func getApps(r *gin.RouterGroup) {
	r.GET("/apps", func(context *gin.Context) {
		if Database.MysqlDB != nil {
			data := Dao.GetApps()

			context.JSON(200, gin.H{
				"apps": data,
			})
		} else {
			context.JSON(500, Error.DatabaseConnectError)
		}
	})
}

func getAppsWithName(r *gin.RouterGroup) {
	r.GET("/apps/ByName", func(context *gin.Context) {
		if Database.MysqlDB != nil {
			name := context.Query("name")
			data := Dao.GetAppsByName(name)

			context.JSON(200, gin.H{
				"apps": data,
			})
		} else {
			context.JSON(500, Error.DatabaseConnectError)
		}
	})
}
