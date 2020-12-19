package Router

import (
	"github.com/gin-gonic/gin"
	"mosk-service/Controller"
)

func SetupRouter(r *gin.Engine) {
	Controller.SetupBaseRouter(r)
}
