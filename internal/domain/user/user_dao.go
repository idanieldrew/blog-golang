package user

import (
	"database/sql"
	"github.com/idanieldrew/blog-golang/internal/repository/postgres"
	"github.com/idanieldrew/blog-golang/pkg/errors/restError"
	"github.com/idanieldrew/blog-golang/pkg/logger"
)

const (
	queryGetUser = "SELECT name,email,phone FROM users WHERE id= $1;"
)

func (u *User) Get() *restError.RestError {
	stmt, err := postgres.Db.Prepare(queryGetUser)
	if err != nil {
		logger.Error("problem in prepare to get user", err)
		return restError.ServerError("server error")
	}

	defer func(stmt *sql.Stmt) {
		if se := stmt.Close(); se != nil {
			return
		}
	}(stmt)

	res := stmt.QueryRow(u.Id)

	if getErr := res.Scan(&u.Name, &u.Email, &u.Phone); getErr != nil {
		logger.Error("problem when try to get user", getErr)
		return restError.ServerError("server error")
	}
	return nil
}
