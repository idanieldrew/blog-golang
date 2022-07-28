package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.New()
)

func StartApp() {
	mapUrls()

	log.Println("Running http:/localhost:8080")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
