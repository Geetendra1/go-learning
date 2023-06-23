package main

import (
	"go-learning/initializers"

	"github.com/RichardKnop/machinery/v1"
	"go.uber.org/zap"
)

var (
	Logger     *zap.SugaredLogger
	taskserver *machinery.Server
)

func init() {
	logger, _ := zap.NewProduction()
	Logger = logger.Sugar()

	taskserver = initializers.GetMachineryServer()

}

// // This function starts a worker for a task server in Go.
func main() {
	Logger.Info("initing worker")

	// This line of code is creating a new worker instance for a task server using the `NewWorker` method
	// provided by the `machinery` package. The first argument `"machinery_worker"` is the name of the
	// worker, and the second argument `10` is the number of concurrent worker processes to run.
	worker := taskserver.NewWorker("machinery_worker", 10)
	if err := worker.Launch(); err != nil {
		Logger.Fatal(err, "could not launch the worker")
	}

}
