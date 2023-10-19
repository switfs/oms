package main

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
	"github.com/switfs/oms/utils"
	"github.com/urfave/cli/v2"
)



func main() {
	_ = logging.SetLogLevel("*", "INFO")

	app := cli.App{
		Commands: []*cli.Command{
			Run,
		},
	}

	app.Setup()
	if err := app.Run(os.Args); err != nil {
		log.Warnf("%+v", err)
		return
	}
}
