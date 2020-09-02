package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ianschenck/envflag"
	"github.com/urfave/cli"

	"github.com/mixj93/react-gin-starter/backend/pkg/server"
)

func main() {
	envflag.Parse()

	app := cli.NewApp()

	app.Name = "react-gin-starter-backend"
	// app.Version = version.Version.String() TODO
	app.Usage = "React Gin Starter Backend"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			EnvVar: "SERVER_ADDR",
			Name:   "server-addr",
			Usage:  "server address",
			Value:  ":5678",
		},
		cli.BoolFlag{
			EnvVar: "DEBUG",
			Name:   "debug",
			Usage:  "start the server in debug mode",
		},
		cli.BoolFlag{
			EnvVar: "DEV",
			Name:   "dev",
			Usage:  "start the server in dev mode",
		},
	}

	app.Action = server.Serve

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func disableGlog() {
	flag.Lookup("logtostderr").Value.Set("true")
}
