package app

import (
	"github.com/idanieldrew/blog-golang/internal/controllers/posts"
	"github.com/idanieldrew/blog-golang/internal/controllers/users"
	"github.com/idanieldrew/blog-golang/internal/middlewares"
)

func mapUrls() {
	authRoute()
	userRoute()
	postRoute()
}

// auth route
func authRoute() {
	auth := router.Group("v1")
	auth.POST("register", users.Register)
	auth.POST("login", users.Login)
}

// user route
func userRoute() {
	user := router.Group("/users")
	//User
	user.GET("/:user_id", users.Get)
}

// post route
func postRoute() {
	post := router.Group("/posts")

	// Set auth middleware
	post.Use(middlewares.AuthMiddleware())

	//Posts
	post.GET("/:post", posts.Get)
	// Store post
	post.POST("/", posts.Store)
}
