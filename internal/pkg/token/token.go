package token

import (
	"github.com/idanieldrew/blog-golang/pkg/errors/restError"
	"github.com/idanieldrew/blog-golang/pkg/logger"
	"github.com/kataras/jwt"
	"time"
)

var sharedKey = jwt.MustGenerateRandom(32)

func GenerateToken(input1, input2 string) (string, *restError.RestError) {
	claims := jwt.Map{
		"email":    input1,
		"password": input2,
	}

	bytes, signErr := jwt.Sign(jwt.HS256, sharedKey, claims, jwt.MaxAge(2*time.Hour))
	if signErr != nil {
		logger.Error("problem when sign", signErr)
		restErr := restError.ServerError("server error")
		return "", restErr
	}

	return string(bytes), nil
}
