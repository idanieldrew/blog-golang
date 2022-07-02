package app

import "github.com/idanieldrew/blog-golang/internal/controllers/users"

func mapUrls() {
	router.GET("/users/:user_id", users.Get)
}