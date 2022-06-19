package repository

import (
	"database/sql"
	"wafflehacks/entities/usertypes"
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
	tx, err := cr.db.Begin()
	if err != nil {
		return nil, &models.ErrorResponse{ErrorMessage: "Не удалось подготовить транзакцию", ErrorCode: 500}
	}
	defer tx.Commit()

	err = tx.QueryRow(`
	INSERT INTO users
		(Firstname, Lastname, username, Email , Password,Age)
	VALUES
		($1,$2,$3,$4,$5,$6)
	RETURNING
		id;`,
		client.Firstname, client.Lastname, client.Username, client.Email, client.Password, client.Age).Scan(&client.ID)
	if err != nil {
		tx.Rollback()
		cr.log.Debug("Не удалось создать психолога по причине: " + err.Error())
		return nil, &models.ErrorResponse{ErrorMessage: "Не удалось зарегстрироваться, попробуйте ввести другой адрес или ник", ErrorCode: 400}
	}

	_, err = tx.Exec(`
	INSERT INTO clients
		(id)
	VALUES
		($1)
	`, client.ID)
	if err != nil {
		tx.Rollback()
		cr.log.Debug("Не удалось создать клиента по причине: " + err.Error())
		return nil, &models.ErrorResponse{ErrorMessage: "Не удалось зарегистрироваться", ErrorCode: 400}
	}

	_, err = tx.Exec(`
	INSERT INTO Usertype
		(email, role)
	VALUES
		($1,$2)
	`, client.Email, usertypes.Client)
	if err != nil {
		tx.Rollback()
		cr.log.Debug("Не удалось создать психолога по причине: " + err.Error())
		return nil, &models.ErrorResponse{ErrorMessage: "Адрес электронной почты занят", ErrorCode: 400}
	}
	return client, nil
}
