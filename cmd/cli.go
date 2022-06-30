package cmd

import (
	"github.com/urfave/cli/v2"
	"os"
)

func Run() error {
	app := cli.App{
		Commands: []*cli.Command{serveCMD},
	}
	return app.Run(os.Args)
}