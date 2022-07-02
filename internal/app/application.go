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

	log.Println("Running")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
