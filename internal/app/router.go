package app

import "github.com/idanieldrew/blog-golang/internal/controllers/users"

func mapUrls() {
	userRoute()
}

func userRoute()  {
	user := router.Group("/users")
	user.GET("/:user_id", users.Get)
}
