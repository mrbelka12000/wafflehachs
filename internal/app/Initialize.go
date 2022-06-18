package app

import (
	"database/sql"
	"wafflehacks/database"
	"wafflehacks/internal/handler"
	"wafflehacks/internal/repository"
	"wafflehacks/internal/service"

	"go.uber.org/zap"
)

func Initialize(log *zap.SugaredLogger) (*handler.Handler, error) {
	db, err := database.GetConnection()
	if err != nil {
		// return nil, err
	}

	// database.UpTables(db, log)
	db = &sql.DB{}
	repo := repository.NewRepo(db, log)
	srv := service.NewService(repo, log)
	return handler.NewHandler(srv, log), nil
}
