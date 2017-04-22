package main

import (
	"github.com/beldpro-ci/reppa/reppa/common"
	"github.com/beldpro-ci/reppa/reppa/core/git"
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
				var gitRoot = c.String("git-root")
				if port == "" {
					return errors.New(
						"`--port` must be specified.")
				}

				if gitRoot == "" {
					return errors.New(
						"`--git-root` must be specified.")
				}

				log.Info("Starting HTTP server",
					zap.String("port", port))

				http.ListenAndServe(":"+port, router.New(&router.RouterConfig{
					GitConfig: &git.GitConfig{
						RepositoriesRoot: gitRoot,
					},
				}))
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "port, p",
					Usage:  "`port` to bind to",
					EnvVar: "REPPA_PORT",
					Value:  "8080",
				},

				cli.StringFlag{
					Name:   "git-root, g",
					Usage:  "root `dir` of git repositories",
					EnvVar: "REPPA_GIT_ROOT",
					Value:  "/tmp/git",
				},
			},
		},
	}
	app.Run(os.Args)
}
