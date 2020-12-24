package Router

import (
	"github.com/gin-gonic/gin"
	"mock-service/Controller"
)

func SetupRouter(r *gin.Engine) {
	Controller.SetupBaseRouter(r)
	Controller.SetupTestRouter(r)
	Controller.SetupSystemRouter(r)
	Controller.SetupNcdaRouter(r)
}
