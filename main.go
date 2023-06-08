package main

import (
	"go-learning/controllers"
	"go-learning/initializers"
	"go-learning/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	// signup
	r.POST("/signup", controllers.Signup)

	// login
	r.POST("/signin", controllers.Login)

	// validate
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// all the posts
	r.GET("/posts", controllers.PostsList)

	// create a post
	r.POST("/posts", controllers.PostsCreate)

	// get single post
	r.GET("/posts/:id", controllers.PostsShow)

	// update the post
	r.PUT("/posts/:id", controllers.PostsUpdate)

	// delete the post
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
