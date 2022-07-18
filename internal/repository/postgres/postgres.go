package postgres

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/idanieldrew/blog-golang/internal/config"
	"github.com/idanieldrew/blog-golang/internal/repository"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
)

type postgres struct {
	db *sql.DB
}

var (
	Db *sql.DB
)

func New(cfg config.Pgsql) (repository.PostgresQL, error) {
	var err error
	Db, err = sql.Open("postgres", dsn(cfg))
	if err != nil {
		return nil, err
	}

	if pe := Db.Ping(); pe != nil {
		return nil, pe
	}

	/*	if me := migration(cfg); me != nil {
		return nil, me
	}*/

	return &postgres{db: Db}, nil
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

func Migration(cfg config.Pgsql) error {
	m, err := migrate.New("file://internal/migrations", cfg.PgUrl)
	if err != nil {
		return err
	}

	if ue := m.Up(); ue != nil {
		return ue
	}
	return nil
}
