package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	conf := loadConfig()

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/base/config", func(c *gin.Context) {
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
