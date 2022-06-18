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

func (cr *ClientRepo) SignUp(client *models.Client) (*models.Client, *models.ErrorResponse) {
	err := cr.db.QueryRow(`
	INSERT INTO Clients
		(Firstname, Lastname, Nickname, Email , Password,Age)
	VALUES
		($1,$2,$3,$4,$5,$6)
	RETURNING
		id;`,
		client.Firstname, client.Lastname, client.Username, client.Email, client.Password, client.Age).Scan(&client.ID)
	if err != nil {
		cr.log.Debug("Не удалось создать психолога по причине: " + err.Error())
		return nil, &models.ErrorResponse{ErrorMessage: "Не удалось зарегстрироваться, попробуйте ввести другой адрес или ник", ErrorCode: 400}
	}


	
	return client, nil
}
