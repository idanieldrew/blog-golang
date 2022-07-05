package post

import (
	"database/sql"
	"github.com/idanieldrew/blog-golang/internal/repository/postgres"
	"github.com/idanieldrew/blog-golang/pkg/errors/restError"
)

const (
	queryGetPost = "SELECT id,title,slug,details,description,blueTick,user_id,created_at,updated_at FROM posts WHERE slug = $1;"
)

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
