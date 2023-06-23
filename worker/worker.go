package worker

import (
	"github.com/RichardKnop/machinery/v1"
	"go.uber.org/zap"
)

var (
	Logger *zap.SugaredLogger
)

func init() {
	logger, _ := zap.NewProduction()
	Logger = logger.Sugar()
}

// This function starts a worker for a task server in Go.
func StartWorker(taskserver *machinery.Server) error {
	Logger.Info("initing worker")

	// This line of code is creating a new worker instance for a task server using the `NewWorker` method
	// provided by the `machinery` package. The first argument `"machinery_worker"` is the name of the
	// worker, and the second argument `10` is the number of concurrent worker processes to run.
	worker := taskserver.NewWorker("machinery_worker", 10)
	if err := worker.Launch(); err != nil {
		return err
	}
	return nil
}
