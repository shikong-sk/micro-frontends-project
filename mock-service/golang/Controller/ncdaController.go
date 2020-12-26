package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mock-service/Proxy"
	"strings"
)

type RouterLinked struct {
	router  string
	handler gin.HandlerFunc
}

const ncdaProxyTarget = "http://192.168.1.241:1080"

var ncdaRouterLinked = make([]RouterLinked, 0, 10)

func getRegionRouter() {
	ncdaRouterLinked = append(ncdaRouterLinked, RouterLinked{
		router: "/ncda/basicinfo/region/get",
		handler: func(context *gin.Context) {
			context.JSON(200, gin.H{
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
		},
	})
}

func deviceFindById() {
	ncdaRouterLinked = append(ncdaRouterLinked, RouterLinked{
		router: "/ncda/base/device/find-by-id",
		handler: func(context *gin.Context) {
			context.JSON(200, gin.H{
				"code": 200,
				"msg":  nil,
				"data": gin.H{
					"id":                   136,
					"address":              "dd",
					"authMethodId":         -1,
					"baiduLatitude":        "22.583908",
					"baiduLongitude":       "113.728650",
					"createTime":           "2018-09-20 09:35:48",
					"deviceCode":           "72439149X00E05C680B22",
					"deviceTypeId":         4,
					"fenceRadius":          0,
					"floor":                "",
					"isAp":                 0,
					"isDataOnline":         0,
					"isFence":              0,
					"isMovable":            0,
					"isServiceOnline":      0,
					"lastDataTime":         "2018-09-21 13:25:52",
					"lastHeartbeatTime":    "2018-09-21 13:25:52",
					"latitude":             "23.080433",
					"longitude":            "113.717616",
					"mac":                  "00-E0-5C-68-0B-22",
					"name":                 "3.54人脸设备",
					"placeCode":            "44060510000001",
					"placeId":              1,
					"qqLatitude":           "23.083908",
					"qqLongitude":          "113.728650",
					"remark":               "",
					"runningStateId":       3,
					"serialNum":            "963622996770",
					"subwayCompartmentNum": "",
					"subwayLineInfo":       "",
					"subwayStation":        "",
					"subwayVehicleInfo":    "",
					"updateTime":           "2018-09-20 09:36:44",
					"uploadTimeInterval":   0,
					"vehicleCode":          "",
					"vendorId":             1,
					"isHotspot":            0,
				},
			})
		},
	})
}

func treeDeviceSearch() {
	ncdaRouterLinked = append(ncdaRouterLinked, RouterLinked{
		router: "/ncda/index/tree-device-search",
		handler: func(context *gin.Context) {
			context.JSON(200, gin.H{
				"code": 200,
				"msg":  nil,
				"data": gin.H{
					"id":    2,
					"label": "佛山市",
					"lat":   nil,
					"lng":   nil,
					"children": []gin.H{
						{
							"id":    18,
							"label": "南海",
							"lat":   nil,
							"lng":   nil,
							"children": []gin.H{
								{
									"id":    19,
									"label": "桂城派出所（演示）",
									"lat":   nil,
									"lng":   nil,
									"children": []gin.H{
										{
											"id":    1,
											"label": "场所001",
											"lat":   nil,
											"lng":   nil,
											"children": []gin.H{
												{
													"id":       136,
													"label":    "3.54人脸设备",
													"lat":      "23.383908",
													"lng":      "113.728650",
													"children": []gin.H{},
													"leaf":     true,
												},
											},
											"leaf": false,
										},
									},
									"leaf": false,
								},
							},
							"leaf": false,
						},
					},
					"leaf": false,
				},
			})
		},
	})
}

func hotspotHoverDialogContent() {
	ncdaRouterLinked = append(ncdaRouterLinked, RouterLinked{
		router: "/ncda/hotspot/hotspot/hover-dialog/content",
		handler: func(context *gin.Context) {
			context.JSON(200, gin.H{
				"code": 200,
				"msg":  "OK",
				"data": gin.H{
					"name":        "3.54人脸设备",
					"deviceCode":  "72439149X00E05C680B22",
					"mac":         "00-E0-5C-68-0B-22",
					"faceLogList": []gin.H{},
					"licenseplateLogList": []gin.H{
						{
							"deviceName":        "3.58人Bor车牌采集-桂就派出所-金色强城大门大大謝猛men-",
							"vendor_code":       "72439149X",
							"place_code":        "44068539718276",
							"device_code":       "72439149X78E14C6840FC",
							"district_code":     "440605",
							"sub_district_code": "440605530000",
							"log_time":          1693421867,
							"entrance_time":     1683087426630,
							"gps_longitude":     113.164883,
							"gps_latitude":      23.847736,
							"baidu_longitude":   113.164683,
							"baidu_latitude":    23.947736,
							"version":           1,
							"index":             0,
							"microsecond":       1603887391628961,
							"fullview-path":     "license_plate/2020/10/19/14/1603087394540027846.fullview.4944.nc2tiny2",
							"plate_num":         "粤4047029",
							"plate_color":       1,
							"plate_type":        1,
						},
					},
				},
			})
		},
	})
}

func SetupNcdaRouter(r *gin.Engine) {
	getRegionRouter()
	deviceFindById()
	treeDeviceSearch()
	hotspotHoverDialogContent()

	fmt.Println(ncdaRouterLinked)
	router := r.Group("/ncda")

	router.GET("/*path", func(context *gin.Context) {
		processProxyRouter(ncdaProxyTarget, context)
	})
}

func processProxyRouter(target string, context *gin.Context) bool {
	for _, routerLinked := range ncdaRouterLinked {
		if strings.HasPrefix(context.Request.RequestURI, routerLinked.router) {
			fmt.Println(strings.HasPrefix(context.Request.RequestURI, routerLinked.router))
			routerLinked.handler(context)
			return true
		}
	}
	Proxy.ServeHttpUrl(context.Writer, context.Request, target)
	return false
}