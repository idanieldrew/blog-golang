package user_service

import (
	"github.com/idanieldrew/blog-golang/internal/entity/model/user"
	"github.com/idanieldrew/blog-golang/pkg/errors/restError"
)

var (
	UserService userServiceInterface = &userService{}
)

type (
	userService struct{}

	userServiceInterface interface {
		GetUser(int64) (*user.User, *restError.RestError)
	}
)

func (s *userService) GetUser(userId int64) (*user.User, *restError.RestError) {
	res := &user.User{
		Id: userId,
	}

	if err := res.Get(); err != nil {
		return nil, err
	}

	return res, nil
}