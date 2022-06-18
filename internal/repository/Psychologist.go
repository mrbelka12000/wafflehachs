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
		(Firstname, Lastname, Username, Email , Password,Age)
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

func (pr *PsychologistRepo) GetAll() ([]models.Psychologist, *models.ErrorResponse) {
	psychos := []models.Psychologist{}
	rows, err := pr.db.Query(`
	SELECT 
		id, firstname, lastname, username, avatarurl, age
	FROM 
		psychologists
	`)
	if err != nil {
		pr.log.Debug("Ошибка получения: " + err.Error())
		return nil, &models.ErrorResponse{ErrorMessage: "Не удалось получить список", ErrorCode: 500}
	}

	for rows.Next() {
		avatarUrl := sql.NullString{}

		psycho := models.Psychologist{}
		if err := rows.Scan(&psycho.ID, &psycho.Firstname, &psycho.Lastname, &psycho.Username, &avatarUrl, &psycho.Age); err != nil {
			pr.log.Debug("Ошибка при получении данных психолога: " + err.Error())
			continue
		}

		var avgRate float64
		err = pr.db.QueryRow(`
		SELECT
			AVG(rate)
		FROM
			psychorate
		WHERE
			psychoID=$1
		`, psycho.ID).Scan(&avgRate)
		if err != nil {
			pr.log.Debug("Не удалось посчитать средний рейтинг психолога")
			continue
		}

		psycho.Rate = avgRate
		psycho.AvatarUrl = avatarUrl.String
		psychos = append(psychos, psycho)
	}
	return psychos, nil
}
