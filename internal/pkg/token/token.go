package token

import (
	"github.com/idanieldrew/blog-golang/pkg/errors/restError"
	"github.com/idanieldrew/blog-golang/pkg/logger"
	"github.com/kataras/jwt"
	"time"
)

type (
	claims struct {
		Email string `json:"email"`
	}

	header struct {
		Kid string `json:"kid"`
		Alg string `json:"alg"`
	}
)

func GenerateToken(input1 string) (string, *restError.RestError) {
	privateKey, loadKeyErr := jwt.LoadPrivateKeyRSA(".keys/rsa_private_key_go.pem")
	if loadKeyErr != nil {
		logger.Error("problem in load rsa key", loadKeyErr)
		restErr := restError.ServerError("server error")
		return "", restErr
	}

	claims := claims{Email: input1}
	header := header{
		Kid: "my_key_id_1",
		Alg: jwt.RS256.Name(),
	}

	bytes, signErr := jwt.SignWithHeader(jwt.RS256, privateKey, claims, header, jwt.MaxAge(1*time.Hour))
	if signErr != nil {
		logger.Error("problem when sign", signErr)
		restErr := restError.ServerError("server error")
		return "", restErr
	}

	return string(bytes), nil
}
