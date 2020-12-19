package Controller

import "github.com/gin-gonic/gin"

func SetupBaseRouter(r *gin.Engine){
	router := r.Group("/base")
	getConfig(router)
}

func getConfig(r *gin.RouterGroup){
	r.GET("/config", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"apps": []gin.H{
				{
					"name":       "vue",
					"entry":      "http://127.0.0.1:10001/",
					"container":  "#vue",
					"activeRule": "/vue",
				},
				{
					"name":       "ncda",
					"entry":      "http://127.0.0.1:10002/ncda/",
					"container":  "#ncda",
					"activeRule": "/ncda",
				},
			},
		})
	})
}