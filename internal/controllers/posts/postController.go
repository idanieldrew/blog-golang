package posts

import (
	"github.com/gin-gonic/gin"
	"github.com/idanieldrew/blog-golang/internal/domain/post"
	"github.com/idanieldrew/blog-golang/internal/services/post_service"
	"github.com/idanieldrew/blog-golang/pkg/errors/restError"
	"net/http"
)

func getPostBySlug(slug string) string {
	return slug
}

// Get post
func Get(ctx *gin.Context) {
	slug := getPostBySlug(ctx.Param("post"))

	post, err := post_service.PostService.Get(slug)
	if err != nil {
		ctx.JSON(err.Status, err)
	}

	ctx.JSON(http.StatusOK, post.Marshal(ctx.GetHeader("X-Public") == "true"))
}

// Store post
func Store(ctx *gin.Context) {
	var p post.Post
	if err := ctx.ShouldBindJSON(&p); err != nil {
		restErr := restError.BadRequest("invalid json")
		ctx.JSON(restErr.Status, restErr)
	}

	res, serviceErr := post_service.PostService.Store(p)
	if serviceErr != nil {
		ctx.JSON(serviceErr.Status, serviceErr)
		return
	}

	ctx.JSON(http.StatusCreated, res.Marshal(true))
}
