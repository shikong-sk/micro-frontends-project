package Controller

import (
	"github.com/gin-gonic/gin"
	"mock-service/Service"
)

func SetupSystemRouter(r *gin.Engine) {
	router := r.Group("/system")
	getNetWork(router)
}

func getNetWork(r *gin.RouterGroup) {
	r.GET("/getNetWork", func(context *gin.Context) {
		data := Service.NetWork()
		context.JSON(200, data)
	})
}
