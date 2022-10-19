package main

import (
	"context"
	"github.com/ppxb/go-fiber/internal/server"
	"github.com/ppxb/go-fiber/pkg/logger"
	"github.com/urfave/cli/v2"
	"os"
)

var VERSION = "1.0.0"

func main() {
	ctx := logger.NewTagContext(context.Background(), "__main__")

	app := cli.NewApp()
	app.Name = "go fiber"
	app.Version = VERSION
	app.Usage = "A RBAC eam system based on Gin."
	app.Commands = []*cli.Command{
		rootCmd(ctx),
	}
	err := app.Run(os.Args)
	if err != nil {
		logger.WithContext(ctx).Errorf(err.Error())
	}
}

func rootCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "Start app server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "mode",
				Aliases: []string{"m"},
				Value:   "debug",
				Usage:   "Set app server mode",
			},
		},
		Action: func(c *cli.Context) error {
			return server.Run(
				ctx,
				server.SetVersion(VERSION),
				server.SetMode(c.String("mode")),
			)
		},
	}
}
