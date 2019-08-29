package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

//AuthRequired grants access to authenticated users, requires SharedData middleware
func LoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := sessions.Default(c)
		_, b := SessionUserId(sess)
		if !b {
			c.Redirect(http.StatusSeeOther, "/user/login")
			c.Abort()
		} else {
			c.Next()
		}
	}
}
