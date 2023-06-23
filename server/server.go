package server

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/gin-gonic/gin"
	// "go-learning/utils"
)

func StartServer(taskserver *machinery.Server) {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.Run()
}
