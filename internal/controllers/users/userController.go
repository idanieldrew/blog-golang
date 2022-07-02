package users

import (
	"github.com/gin-gonic/gin"
	"github.com/idanieldrew/blog-golang/internal/services/user_service"
	"github.com/idanieldrew/blog-golang/pkg/errors/restError"
	"net/http"
	"strconv"
)

func getUserById(id string) (int64, *restError.RestError) {
	userId, userErr := strconv.ParseInt(id, 10, 64)
	if userErr != nil {
		return 0, restError.BadRequest("user_service id should be a number")
	}
	return userId, nil
}

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
