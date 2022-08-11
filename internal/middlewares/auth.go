package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/idanieldrew/blog-golang/internal/pkg/token"
	"github.com/idanieldrew/blog-golang/pkg/errors/restError"
	"github.com/idanieldrew/blog-golang/pkg/logger"
	"github.com/kataras/jwt"
	"log"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := context.Request.Header.Get("Authorization")

		successVrf, verifyErr := jwt.VerifyWithHeaderValidator(jwt.RS256, token.Keys, []byte(auth), token.ValidateHeader)
		if verifyErr != nil {
			logger.Error("token not valid", verifyErr)
			restError.UnauthorizedError("token not valid")
		}

		var c token.Claims
		err := successVrf.Claims(&c)
		if err != nil {
			log.Fatalln(err)
		}
		context.Next()
	}
}
