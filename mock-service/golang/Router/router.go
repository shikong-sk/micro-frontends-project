package Router

import (
	"github.com/gin-gonic/gin"
	"mosk-service/Controller"
)

func SetupRouter(r *gin.Engine) *gin.Engine {
	Controller.SetupBaseRouter(r)
	return r
}
