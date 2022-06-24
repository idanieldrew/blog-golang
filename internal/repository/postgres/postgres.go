package postgres

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/idanieldrew/blog-golang/internal/config"
	"github.com/idanieldrew/blog-golang/internal/repository"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
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

	if me := migration(cfg); me != nil {
		return nil, me
	}

/*	driver, ie := postgres2.WithInstance(db, &postgres2.Config{})
	if ie != nil {
		return nil, ie
	}

	m, me := migrate.NewWithDatabaseInstance(
		"file///migrations",
		"weblog",
		driver,
	)
	if me != nil {
		return nil, me
	}

	_ = m.Up()*/
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

func migration(cfg config.Pgsql) error {
	m, err := migrate.New("file:///migrations", "postgres://daniel:secret@localhost:5432/database?sslmode=disable")
	if err != nil {
		return err
	}

	if ue := m.Up(); ue != nil {
		return ue
	}

	return nil

}
