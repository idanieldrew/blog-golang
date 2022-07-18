package post

import (
	"database/sql"
	"github.com/idanieldrew/blog-golang/internal/repository/postgres"
	"github.com/idanieldrew/blog-golang/pkg/errors/restError"
	"github.com/idanieldrew/blog-golang/pkg/logger"
)

const (
	queryGetPost   = "SELECT id,title,slug,details,description,blueTick,user_id,created_at,updated_at FROM posts WHERE slug = $1;"
	queryStorePost = "INSERT INTO posts(title,slug,details,description) VALUES($1,$2,$3,$4)"
)

// Query for get post
func (p *Post) Get() *restError.RestError {
	stmt, err := postgres.Db.Prepare(queryGetPost)

	if err != nil {
		return restError.ServerError("server error")
	}

	defer func(stmt *sql.Stmt) {
		if se := stmt.Close(); se != nil {
			return
		}
	}(stmt)

	row := stmt.QueryRow(p.Slug)
	if re := row.Scan(&p.Id, &p.Title, &p.Slug, &p.Details, &p.Description, &p.BlueTick, &p.UserId, &p.CreatedAt, &p.UpdatedAt); re != nil {
		return restError.ServerError("server error")
	}

	return nil
}

// Query for store post
func (p *Post) Save() *restError.RestError {
	stmt, err := postgres.Db.Prepare(queryStorePost)
	if err != nil {
		logger.Error("problem in prepare save query post", err)
		return restError.ServerError("db error")
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			return
		}
	}(stmt)

	row, qe := stmt.Exec(p.Title, p.Slug, p.Details, p.Description)
	if qe != nil {
		logger.Error("problem when trying to save post", qe)
		return restError.ServerError("db error")
	}

	id, idErr := row.LastInsertId()
	if idErr != nil {
		logger.Error("problem when trying to last post id", qe)
		return restError.ServerError("db error")
	}
	p.Id = int(id)
	return nil
}
