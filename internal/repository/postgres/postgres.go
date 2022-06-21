package postgres

import (
	"database/sql"
	"fmt"
	"github.com/idanieldrew/blog-golang/internal/config"
	"github.com/idanieldrew/blog-golang/internal/repository"
	_ "github.com/lib/pq"
)

type postgres struct {
	db *sql.DB
}

func New(cfg config.Pgsql) (repository.PostgresQL, error) {
	db, err := sql.Open("postgres", dsn(cfg))
	if err != nil {
		return nil, err
	}

	if pe := db.Ping(); pe != nil {
		return nil, pe
	}

	return &postgres{db: db}, nil
}

func dsn(cfg config.Pgsql) string {
	return fmt.Sprintf(
		"user=%v password=%v dbname=%v port=%v",
		cfg.Username,
		cfg.Password,
		cfg.Dbname,
		cfg.Port,
	)
}
