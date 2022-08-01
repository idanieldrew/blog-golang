package users

import (
	"github.com/gin-gonic/gin"
	"github.com/idanieldrew/blog-golang/internal/pkg/token"
	"github.com/idanieldrew/blog-golang/internal/request"
	"github.com/idanieldrew/blog-golang/internal/services/user_service"
	"github.com/idanieldrew/blog-golang/pkg/errors/restError"
	"net/http"
	"strconv"
)

// return id
func getUserById(id string) (int64, *restError.RestError) {
	userId, userErr := strconv.ParseInt(id, 10, 64)
	if userErr != nil {
		return 0, restError.BadRequestError("user_service id should be a number")
	}
	return userId, nil
}

// Get users
func Get(ctx *gin.Context) {
	userId, idErr := getUserById(ctx.Param("user_id"))
	if idErr != nil {
		ctx.JSON(idErr.Status, idErr)
		return
	}

	user, getError := user_service.UserService.GetUser(userId)
	if getError != nil {
		ctx.JSON(getError.Status, getError)
		return
	}

	ctx.JSON(http.StatusOK, user.Marshal(ctx.GetHeader("X-Public") == "true"))
}

// Register user
func Register(ctx *gin.Context) {
	// validation data
	var input request.RegisterUser
	if err := ctx.ShouldBindJSON(&input); err != nil {
		restErr := restError.ValidationError(err.Error())
		ctx.JSON(restErr.Status, restErr)
		return
	}
	user, creatErr := user_service.UserService.Create(input)
	if creatErr != nil {
		ctx.JSON(creatErr.Status, creatErr)
		return
	}
	ctx.JSON(http.StatusOK, user.Marshal(false))
}

// Login user
func Login(ctx *gin.Context) {
	var input request.LoginUser
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		restErr := restError.ValidationError(err.Error())
		ctx.JSON(restErr.Status, restErr)
		return
	}

	user, findErr := user_service.UserService.FindUser(input)
	if findErr != nil {
		ctx.JSON(findErr.Status, findErr)
		return
	}
	generateToken, generateErr := token.GenerateToken(input.Email)
	if generateErr != nil {
		ctx.JSON(generateErr.Status, generateErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":     user.Marshal(false),
		"token":    generateToken,
		"response": 200,
	})

}
