package controllers

import (
	"go-learning/initializers"
	"go-learning/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get data from req
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsList(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find((&posts))

	// return
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostsShow(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	//get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// get data from req
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// find the data
	var post models.Post
	initializers.DB.First(&post, id)

	// update the data
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// find the data
	var post models.Post
	initializers.DB.Delete(&post, id)

	// return
	c.JSON(200, gin.H{
		"message": "deleted successfully",
	})
	// c.Status(200)
}
