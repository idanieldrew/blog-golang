package cmd

import (
	"github.com/idanieldrew/blog-golang/internal/app"
	"github.com/idanieldrew/blog-golang/internal/config"
	"github.com/idanieldrew/blog-golang/internal/repository/postgres"
	"github.com/urfave/cli/v2"
	"log"
)

const path = "build/config/config.yml"

var (
	serveCMD = &cli.Command{
		Name:    "serve",
		Aliases: []string{"serve"},
		Usage:   "serve http",
		Action:  serve,
	}

	migrateCMD = &cli.Command{
		Name:        "migrate",
		Aliases:     []string{"migrate"},
		Usage:       "migration",
		Description: "migration database",
		Category:    "Db",
		Action:      migrate,
	}
)

func serving() (*config.Config, error) {
	cfg := new(config.Config)
	if err := config.Read(path, cfg); err != nil {
		return nil, err
	}

	_, err := postgres.New(cfg.Postgres)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// Serving project
func serve(c *cli.Context) error {
	_, err := serving()
	if err != nil {
		return err
	}
	app.StartApp()

	return nil
}

// Migrate
func migrate(c *cli.Context) error {
	cfg, err := serving()
	if err != nil {
		return err
	}
	if me := postgres.Migration(cfg.Postgres, c.Args().Get(0) == "fresh"); me != nil {
		return me
	}
	log.Print("Successfully migrate")
	return nil
}
