package main

import (
	"go-learning/initializers"
	"go-learning/worker"
	"os"

	"github.com/RichardKnop/machinery/v1"
	"github.com/urfave/cli"
)

var (
	app        *cli.App
	taskserver *machinery.Server
)

func init() {
	// initialize App
	app = cli.NewApp()
	app.Name = "go-machinery"
	app.Usage = "machinery worker and send example tasks with machinery send"
	app.Author = "Geetendra Singh Sengar"
	app.Email = "sengargeetendra123@gmail.com"

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	taskserver = initializers.GetMachineryServer()
	// worker.StartWorker(taskserver)
}

func main() {

	app.Commands = []cli.Command{
		// {
		// 	Name:  "server",
		// 	Usage: "Run the server that takes task input",
		// 	Action: func(c *cli.Context) {
		// 		initializers.Logger.Info("server")
		// 		initializers.StartServer(taskserver)
		// 	},
		// },
		{
			Name:  "worker",
			Usage: "Run the worker that will consume tasks",
			Action: func(c *cli.Context) {
				initializers.Logger.Info("worker")
				worker.StartWorker(taskserver)
			},
		},
	}

	app.Run(os.Args)

}

// func main() {

// 	app.Commands = []cli.Command{
// 		{
// 			Name:  "worker",
// 			Usage: "Run the worker that will consume tasks",
// 			Action: func(c *cli.Context) {
// 				initializers.Logger.Info("worker")
// 				worker.StartWorker(taskserver)
// 			},
// 		},
// 	}

// 	r := gin.Default()

// 	r.GET("/", func(ctx *gin.Context) {
// 		ctx.JSON(200, gin.H{
// 			"message": "Hello World",
// 		})
// 	})

// 	// signup
// 	r.POST("/signup", controllers.Signup)

// 	// login
// 	r.POST("/signin", controllers.Login)

// 	// validate
// 	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

// 	// all the posts
// 	r.GET("/posts", controllers.PostsList)

// 	// create a post
// 	r.POST("/posts", controllers.PostsCreate)

// 	// get single post
// 	r.GET("/posts/:id", controllers.PostsShow)

// 	// update the post
// 	r.PUT("/posts/:id", controllers.PostsUpdate)

// 	// delete the post
// 	r.DELETE("/posts/:id", controllers.PostsDelete)

// 	app.Run(os.Args)
// 	r.Run()
// }
