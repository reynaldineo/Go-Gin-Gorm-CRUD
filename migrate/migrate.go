package main

import (
	"github.com/reynaldineo/Go-Gin-Gorm-CRUD/initializers"
	"github.com/reynaldineo/Go-Gin-Gorm-CRUD/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}