package main

import (
	"github.com/beldpro-ci/reppa/reppa/common"
	"github.com/beldpro-ci/reppa/reppa/router"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gopkg.in/urfave/cli.v1"
	"net/http"
	"os"
)

var log = common.GetLogger()

func main() {
	app := cli.NewApp()
	app.Name = "reppa"
	app.Usage = "Repositories Manager"
	app.Commands = []cli.Command{
		{
			Name: "start",
			Action: func(c *cli.Context) error {
				var port = c.String("port")
				if port == "" {
					return errors.New(
						"`--port` must be specified.")
				}

				log.Info("Starting HTTP server",
					zap.String("port", port))

				http.ListenAndServe(":"+port, router.New())
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "port",
					EnvVar: "REPPA_PORT",
					Value:  "8080",
				},
			},
		},
	}
	app.Run(os.Args)
}
