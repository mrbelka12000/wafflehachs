package main

import (
	"os"
	"wafflehacks/internal/app"
	"wafflehacks/internal/server"

	"go.uber.org/zap"
)

func main() {
	// tools.Loadenv()
	logger, err := zap.NewDevelopment()
	if err != nil {
		return
	}
	log := logger.Sugar()
	defer log.Sync()
	handler, err := app.Initialize(log)
	if err != nil {
		log.Debug("Error while conecting to postgres: ", err)
	}
	srv := server.NewServer(handler)
	log.Info("Server started on port: " + os.Getenv("Port"))
	log.Fatal(srv.ListenAndServe())
}
