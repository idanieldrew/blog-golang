package posts

import (
	"github.com/gin-gonic/gin"
	"github.com/idanieldrew/blog-golang/internal/services/post_service"
	"net/http"
)

func getPostBySlug(slug string) string {
	return slug
}

func Get(ctx *gin.Context) {
	slug := getPostBySlug(ctx.Param("post"))

	post, err := post_service.PostService.Get(slug)
	if err != nil {
		ctx.JSON(err.Status, err)
	}

	ctx.JSON(http.StatusOK, post.Marshal(ctx.GetHeader("X-Public") == "true"))
}