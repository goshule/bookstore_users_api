package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	log.Println("Application start request")
	map_Urls()
	router.Run(":8080")
}
