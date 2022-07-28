package user

import (
	"database/sql"
	"github.com/idanieldrew/blog-golang/internal/repository/postgres"
	"github.com/idanieldrew/blog-golang/pkg/errors/restError"
	"github.com/idanieldrew/blog-golang/pkg/logger"
)

const (
	queryGetUser    = "SELECT name,email,phone FROM users WHERE id= $1;"
	queryInsertUser = "INSERT INTO users(name,email,phone,password,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id"
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

func (u *User) Store() *restError.RestError {
	lastId := 0
	if err := postgres.Db.QueryRow(queryInsertUser,
		u.Name, u.Email, u.Phone, u.Password, u.CreatedAt, u.UpdatedAt).
		Scan(&lastId); err != nil {
		logger.Error("problem in query row to insert user", err)
		return restError.ServerError("server error")
	}

	u.Id = int64(lastId)
	return nil
}
