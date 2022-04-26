package main

import (
	"Golang-API-tutoriel/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	controllers.InitDatabase()

	r := gin.Default()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
