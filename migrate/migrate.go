package main

import (
	"go-learning/initializers"
	"go-learning/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	// initializers.DB.AutoMigrate(&models.User{})
}
