package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sellimenes/shadecom-backend/initializers"
	"github.com/sellimenes/shadecom-backend/models"
)

func PostCreate(c *gin.Context) {
// Get data off req body
var body struct {
	Body string
	Title string
}

c.Bind(&body)

// Create a post
post := models.Post{Title: body.Title, Body: body.Body}
result := initializers.DB.Create(&post)

if result.Error != nil {
	c.Status(400)
	return
}

// Return post

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	// Get all posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Return all posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	// Get post by id
	var post models.Post
	initializers.DB.First(&post, c.Param("id"))

	// Return post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get post by id
	var post models.Post
	initializers.DB.First(&post, c.Param("id"))

	// Get data off req body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	// Update post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, 
		Body: body.Body,
	})

	// Return post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// Get post by id
	var post models.Post
	initializers.DB.First(&post, c.Param("id"))

	// Delete post
	initializers.DB.Delete(&post)

	// Return post
	c.JSON(200, gin.H{
		"post": post,
	})
}