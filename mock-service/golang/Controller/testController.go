package Controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetupTestRouter(r *gin.Engine) {
	router := r.Group("/test")
	getSession(router)
}

func getSession(r *gin.RouterGroup) {
	r.GET("/getSession", func(context *gin.Context) {
		session := sessions.Default(context)
		session.Set("test","sessionTest")
		_ = session.Save()
		context.JSON(200, session.Get("test"))
	})
}