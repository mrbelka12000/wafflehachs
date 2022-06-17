package main

import (
	"fmt"
	"wafflehacks/database"
	"wafflehacks/tools"

	"go.uber.org/zap"
)

func main() {
	tools.Loadenv()
	logger, err := zap.NewDevelopment()
	if err != nil {
		return
	}
	sugar := logger.Sugar()
	//nolint
	defer sugar.Sync()

	db, err := database.GetConnection()
	if err != nil {
		sugar.Fatalf(err.Error())
	}
	fmt.Println("starting wafflehacks", db)
}
