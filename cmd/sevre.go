package cmd

import (
	"github.com/idanieldrew/blog-golang/internal/config"
	"github.com/idanieldrew/blog-golang/internal/repository/postgres"
	"github.com/urfave/cli/v2"
)

const path = "build/config/config.yml"

var serveCMD = &cli.Command{
	Name:    "serve",
	Aliases: []string{"s"},
	Usage:   "serve http",
	Action:  serve,
}

func serve(c *cli.Context) error {
	cfg := new(config.Config)
	if err := config.Read(path, cfg); err != nil {
		return err
	}

	_, err := postgres.New(cfg.Postgres)
	if err != nil {
		return err
	}

	// Http server
	//handler := gin.New()

	return nil
}
