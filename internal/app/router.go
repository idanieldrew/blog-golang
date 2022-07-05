package app

import (
	"github.com/idanieldrew/blog-golang/internal/controllers/posts"
	"github.com/idanieldrew/blog-golang/internal/controllers/users"
)

func mapUrls() {
	userRoute()
	postRoute()
}

func userRoute() {
	user := router.Group("/users")
	// Get user
	user.GET("/:user_id", users.Get)
}

func postRoute() {
	post := router.Group("/posts")
	//Get posts
	post.GET("/:post", posts.Get)
}
