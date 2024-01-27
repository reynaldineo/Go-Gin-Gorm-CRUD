package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reynaldineo/Go-Gin-Gorm-CRUD/controllers"
	"github.com/reynaldineo/Go-Gin-Gorm-CRUD/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	router.POST("/posts", controllers.PostsCreate)
	router.GET("/posts", controllers.PostIndex)
	router.GET("/posts/:id", controllers.PostShow)
	router.PUT("/posts/:id", controllers.PostUpdate)
	router.DELETE("/posts/:id", controllers.PostDelete)


	router.Run()
}