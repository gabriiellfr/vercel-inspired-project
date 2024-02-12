// app/main.go

package main

import (
	"log"
	"vercel-inspired-project-v2/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", handlers.IndexHandler)
	router.GET("/build", handlers.BuildHandler)

	log.Fatal(router.Run(":8080"))
}
