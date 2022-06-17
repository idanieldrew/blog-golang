package cmd

import (
	"github.com/idanieldrew/blog-golang/internal/config"
	"github.com/idanieldrew/blog-golang/internal/repository/postgres"
)

func Serve(cfg *config.Config) error {
	_, err := postgres.New(cfg.Postgres)
	if err != nil {
		return err
	}

	return nil
}
