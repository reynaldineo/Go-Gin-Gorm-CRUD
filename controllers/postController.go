package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/reynaldineo/Go-Gin-Gorm-CRUD/initializers"
	"github.com/reynaldineo/Go-Gin-Gorm-CRUD/models"
)

func PostsCreate(c *gin.Context) {
	// GET data from request body
	var body struct{
		Body string
		Title string
	}

	c.Bind(&body) // request body will store in body variable

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post) 

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Error",
		})
		return
	}

	// Return the post
	c.JSON(201, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context){
	// Get the post
	var posts []models.Post 
	initializers.DB.Find(&posts) // SELECT * FROM posts

	// return the post
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context){
	// Get id from url
	id := c.Param("id")

	// Get the post
	var post models.Post 
	initializers.DB.First(&post, id) // get post data by id

	// Return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context){
	// Get id from url
	id := c.Param("id")

	// Get data from request body
	var body struct{
		Body string
		Title string
	}
	c.Bind(&body) // request body will store in body variable

	// Find the post we want to update
	var post models.Post 
	initializers.DB.First(&post, id) // SELECT * FROM posts

	// Update the post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, 
		Body: body.Body,
	})

	// Return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context){
	// Get id from url
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{} , id)

	c.Status(200)
}