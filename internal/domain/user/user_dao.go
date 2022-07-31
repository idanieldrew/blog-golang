package user

import (
	"database/sql"
	"github.com/idanieldrew/blog-golang/internal/repository/postgres"
	"github.com/idanieldrew/blog-golang/pkg/errors/restError"
	"github.com/idanieldrew/blog-golang/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

const (
	queryGetUser                = "SELECT name,email,phone FROM users WHERE id= $1;"
	queryInsertUser             = "INSERT INTO users(name,email,phone,password,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id;"
	queryFindByEmailAndPassword = "SELECT name,phone,password FROM users WHERE email= $1;"
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

func (u *User) FindWithLogin() *restError.RestError {
	password := u.Password

	stmt, err := postgres.Db.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("problem in prepare to get user", err)
		return restError.ServerError("server error")
	}

	defer func(stmt *sql.Stmt) {
		if se := stmt.Close(); se != nil {
			return
		}
	}(stmt)

	row := stmt.QueryRow(u.Email)
	scanErr := row.Scan(&u.Name, &u.Phone, &u.Password)

	// Change correct password
	passErr := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if scanErr != nil || passErr != nil {
		logger.Error("email or pass is incorrect", scanErr)
		return restError.UnauthorizedError("email or pass is incorrect")
	}

	return nil
}
