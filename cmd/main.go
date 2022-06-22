package main

import (
	"go.uber.org/zap"
	"wafflehacks/internal/app"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		println(err.Error())
		return
	}

	log := logger.Sugar()
	defer log.Sync()

	app.Run(log)

	log.Info("Server Exited Properly")
}
