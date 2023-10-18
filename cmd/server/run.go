package main

import (
	"github.com/urfave/cli/v2"
)

var Run = &cli.Command{
	Name:  "run",
	Usage: "运行服务",
	Action: func(cctx *cli.Context) error {
		log.Info("ok")
		return nil
	},
}
