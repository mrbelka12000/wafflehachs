package repository

import (
	"database/sql"
	"wafflehacks/models"

	"go.uber.org/zap"
)

type ClientRepo struct {
	db  *sql.DB
	log *zap.SugaredLogger
}

func newClient(db *sql.DB, log *zap.SugaredLogger) *ClientRepo {
	return &ClientRepo{db, log}
}

func (cr *ClientRepo) SignUp(clientId int) *models.ErrorResponse {

	_, err := cr.db.Exec(`
	INSERT INTO clients
		(id)
	VALUES
		($1)
	`, clientId)
	if err != nil {
		cr.log.Debug("Не удалось создать клиента по причине: " + err.Error())
		return &models.ErrorResponse{ErrorMessage: "Не удалось зарегистрироваться", ErrorCode: 400}
	}

	return nil
}
