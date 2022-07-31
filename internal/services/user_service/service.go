package user_service

import (
	"github.com/idanieldrew/blog-golang/internal/domain/user"
	"github.com/idanieldrew/blog-golang/internal/request"
	"github.com/idanieldrew/blog-golang/pkg/errors/restError"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	UserService userServiceInterface = &userService{}
)

type (
	userService struct{}

	userServiceInterface interface {
		GetUser(int64) (*user.User, *restError.RestError)
		Create(request.RegisterUser) (*user.User, *restError.RestError)
		FindUser(request.LoginUser) (*user.User, *restError.RestError)
	}
)

func (s *userService) FindUser(loginUser request.LoginUser) (*user.User, *restError.RestError) {
	data := &user.User{
		Email:    loginUser.Email,
		Password: loginUser.Password,
	}

	if findErr := data.FindWithLogin(); findErr != nil {
		return nil, findErr
	}

	return data, nil
}

func (s *userService) Create(req request.RegisterUser) (*user.User, *restError.RestError) {
	b, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 8)

	res := &user.User{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  string(b),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		//Type:      1,
	}

	if err := res.Store(); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *userService) GetUser(userId int64) (*user.User, *restError.RestError) {
	res := &user.User{
		Id: userId,
	}

	if err := res.Get(); err != nil {
		return nil, err
	}

	return res, nil
}
