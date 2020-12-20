package Session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const (
	CookieKey = "session_id"
)

var Store cookie.Store

func InitSession(r *gin.Engine, sessionSecret string) {
	Store = cookie.NewStore([]byte(sessionSecret))
	r.Use(sessions.Sessions(CookieKey, Store))
	option := sessions.Options{MaxAge: 1800}
	Store.Options(option)
}
