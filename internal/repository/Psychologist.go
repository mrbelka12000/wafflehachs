package repository

import (
	"database/sql"
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

func (pr *PsychologistRepo) SignUp(psychoId int) *models.ErrorResponse {

	_, err := pr.db.Exec(`
	INSERT INTO Psychologists
		(id)
	VALUES
		($1)
	`, psychoId)
	if err != nil {
		pr.log.Debug("Не удалось создать психолога по причине: " + err.Error())
		return &models.ErrorResponse{ErrorMessage: "Не удалось зарегистрироваться", ErrorCode: 400}
	}

	return nil
}

func (pr *PsychologistRepo) GetAll() ([]models.Psychologist, *models.ErrorResponse) {
	psychos := []models.Psychologist{}
	rows, err := pr.db.Query(`
	SELECT 
		users.id, users.firstname, users.lastname, users.username, users.avatarurl, users.age
	FROM 
		Psychologists 
	JOIN
		users on users.id = Psychologists.id
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
		}

		psycho.Rate = avgRate
		psycho.AvatarUrl = avatarUrl.String
		psychos = append(psychos, psycho)
	}
	return psychos, nil
}
