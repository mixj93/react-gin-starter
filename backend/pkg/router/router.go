package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r gin.IRouter) {
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}
