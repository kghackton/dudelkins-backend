package main

import (
	"context"
	"dudelkins/infrastructure"
	"dudelkins/internal/environment"
	"dudelkins/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

var (
	env environment.Environment
	ctx = context.Background()

	version, date    string
	deploymentEnvTag string
)

func init() {
	if depEnv := os.Getenv("DEPLOYMENT_ENVIRONMENT"); depEnv != "" {
		deploymentEnvTag = depEnv
	} else {
		deploymentEnvTag = "dev"
	}

	moscowTimezone := time.FixedZone("UTC+3", 60*60*3)
	time.Local = moscowTimezone

	env = environment.NewEnvironment()
	logger.Infof("environment: %+v", env)
}

func main() {
	kernel, err := infrastructure.Inject(ctx, env)
	if err != nil {
		logger.Fatal(err)
	}

	var wg sync.WaitGroup
	cancelCtx, cancelFunc := signal.NotifyContext(ctx, os.Kill, os.Interrupt)
	defer cancelFunc()

	// TODO: move port to cfg
	routerStartFunc := initRouter(kernel, "18754")
	wg.Add(1)
	if err := routerStartFunc(&wg); err != nil {
		if err != http.ErrServerClosed {
			logger.Error(err.Error())
		} else {
			logger.Info("server gracefully stopped")
		}
	}

	<-cancelCtx.Done()

	if err := kernel.DbHandler.Close(); err != nil {
		logger.Error(err.Error())
	}

	wg.Wait()
}
