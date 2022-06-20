package repository

import (
	"database/sql"
	"wafflehacks/entities/busymode"
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
		(id, busymode)
	VALUES
		($1, $2)
	`, psychoId, busymode.ActiveMode)
	if err != nil {
		pr.log.Error("Не удалось создать психолога по причине: " + err.Error())
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
		pr.log.Error("Ошибка получения: " + err.Error())
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

func (pr *PsychologistRepo) GetByUsername(username string) (*models.Psychologist, *models.ErrorResponse) {
	psycho := &models.Psychologist{}
	avatar := sql.NullString{}
	err := pr.db.QueryRow(`
	SELECT 
		users.id, users.firstname, users.lastname, users.username,users.email, users.avatarurl, users.age, psychologists.busymode
	FROM 
		Psychologists 
	JOIN
		users on users.id = Psychologists.id
	WHERE
	    users.username=$1
`, username).Scan(&psycho.ID, &psycho.Firstname, &psycho.Lastname, &psycho.Username, &psycho.Email, &avatar, &psycho.Age, &psycho.BusyMode)
	if err != nil {
		pr.log.Error("не удалось получить психлога по причине: " + err.Error())
		return nil, &models.ErrorResponse{ErrorMessage: "Не удалось найти", ErrorCode: 400}
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

	rows, err := pr.db.Query(`
	SELECT 
	    users.id, users.username, psychorate.rate, psychorate.anonym, psychorate.comment
	FROM
	    psychorate
	JOIN 
		users on users.id = psychorate.clientid
	WHERE 
	    psychorate.psychoID=$1
`, psycho.ID)
	if err != nil {
		pr.log.Debug(err.Error())
	}

	reviews := []models.Review{}

	for rows.Next() {
		comment := sql.NullString{}
		review := models.Review{}
		if err = rows.Scan(&review.ID, &review.Username, &review.Rating, &review.Anonym, &comment); err != nil {
			pr.log.Debug(err.Error())
			continue
		}

		if review.Anonym {
			review.Username = "Anonymous"
			review.ID = 0
		}
		review.Comment = comment.String
		reviews = append(reviews, review)
	}

	psycho.Rate = avgRate
	psycho.AvatarUrl = avatar.String
	psycho.Reviews = reviews
	return psycho, nil
}
