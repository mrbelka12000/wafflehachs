package repository

import (
	"database/sql"
	"wafflehacks/entities/usertypes"
	"wafflehacks/models"

	"go.uber.org/zap"
)

type PsychologistRepo struct {
	db  *sql.DB
	log *zap.SugaredLogger
}

func newPsycho(db *sql.DB, log *zap.SugaredLogger) *PsychologistRepo {
	return &PsychologistRepo{db, log}
}

func (pr *PsychologistRepo) SignUp(psycho *models.Psychologist) (*models.Psychologist, *models.ErrorResponse) {
	tx, err := pr.db.Begin()
	if err != nil {
		return nil, &models.ErrorResponse{ErrorMessage: "Не удалось подготовить транзакцию", ErrorCode: 500}
	}
	defer tx.Commit()

	err = tx.QueryRow(`
	INSERT INTO Psychologists
		(Firstname, Lastname, Nickname, Email , Password,Age)
	VALUES
		($1,$2,$3,$4,$5,$6)
	RETURNING
		id;`,
		psycho.Firstname, psycho.Lastname, psycho.Username, psycho.Email, psycho.Password, psycho.Age).Scan(&psycho.ID)
	if err != nil {
		tx.Rollback()
		pr.log.Debug("Не удалось создать психолога по причине: " + err.Error())
		return nil, &models.ErrorResponse{ErrorMessage: "Не удалось зарегстрироваться, попробуйте ввести другой адрес или ник", ErrorCode: 400}
	}

	_, err = tx.Exec(`
	INSERT INTO Usertype
		(email, role)
	VALUES
		($1,$2)
	`, psycho.Email, usertypes.Psycho)
	if err != nil {
		tx.Rollback()
		pr.log.Debug("Не удалось создать психолога по причине: " + err.Error())
		return nil, &models.ErrorResponse{ErrorMessage: "Адрес электронной почты занят", ErrorCode: 400}
	}

	return psycho, nil
}
