package main

import (
	"os"
	"wafflehacks/internal/app"
	"wafflehacks/internal/server"
	"wafflehacks/tools"

	"go.uber.org/zap"
)

func main() {
	tools.Loadenv()
	logger, err := zap.NewDevelopment()
	if err != nil {
		return
	}
	log := logger.Sugar()
	defer log.Sync()
	handler, err := app.Initialize(log)
	if err != nil {
		log.Fatal("Error while conecting to postgres: ", err)
	}
	srv := server.NewServer(handler)
	log.Info("Server started on port: " + os.Getenv("PORT"))
	log.Fatal(srv.ListenAndServe())
}
