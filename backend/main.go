package main

import (
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Serve Statics
	ex, _ := os.Executable()
	dir := path.Dir(ex)
	r.StaticFile("/", dir+"/dist/index.html")
	r.Static("/static", dir+"/dist/static")

	r.Run(":5678") // listen and serve on 0.0.0.0:5678
}
