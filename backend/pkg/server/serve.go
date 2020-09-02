package server

import (
	"os"
	"path"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/mixj93/react-gin-starter/backend/pkg/middlewares"
	"github.com/mixj93/react-gin-starter/backend/pkg/router"
)

func Serve(c *cli.Context) error {
	// Set Log Level
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	// Gin
	r := gin.Default()
	r.Use(gin.Recovery())
	r.UseRawPath = true

	// Register Router
	router.RegisterRouter(r)

	// Serve Statics
	if !c.Bool("dev") {
		ex, _ := os.Executable()
		dir := path.Dir(ex)
		r.LoadHTMLFiles(dir + "/dist/index.html")
		r.Use(static.Serve("/", static.LocalFile(dir+"/dist/", false)))
		r.Use(middlewares.HistoryAPIFallback())
	}

	// Listen and Server Addr, default in 0.0.0.0:5678
	addr := c.String("server-addr")
	r.Run(addr)

	return nil
}
