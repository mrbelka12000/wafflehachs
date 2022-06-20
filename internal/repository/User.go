package repository

import (
	"database/sql"
	"wafflehacks/entities/usertypes"
	"wafflehacks/models"

	"go.uber.org/zap"
)

type UserRepo struct {
	db  *sql.DB
	log *zap.SugaredLogger
}

// ....
func newUser(db *sql.DB, log *zap.SugaredLogger) *UserRepo {
	return &UserRepo{
		db:  db,
		log: log,
	}
}

func (ur *UserRepo) CanLogin(user *models.User) (*models.User, *models.ErrorResponse) {
	User := &models.User{}

	err := ur.db.QueryRow(`
		SELECT 
			id, Password
		FROM
			users
		WHERE
			Email = $1	
		`, user.Email).Scan(&User.ID, &User.Password)
	if err != nil {
		ur.log.Debug("email not found : " + err.Error())
		return nil, &models.ErrorResponse{ErrorMessage: "getUser failed", ErrorCode: 500}
	}

	return User, nil
}

func (ur *UserRepo) SignUp(user *models.User) (*models.User, *models.ErrorResponse) {
	tx, err := ur.db.Begin()
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
		user.Firstname, user.Lastname, user.Username, user.Email, user.Password, user.Age).Scan(&user.ID)
	if err != nil {
		tx.Rollback()
		ur.log.Debug("Не удалось создать психолога по причине: " + err.Error())
		return nil, &models.ErrorResponse{ErrorMessage: "Не удалось зарегстрироваться, попробуйте ввести другой адрес или ник", ErrorCode: 400}
	}

	_, err = tx.Exec(`
	INSERT INTO Usertype
		(email, role)
	VALUES
		($1,$2)
	`, user.Email, usertypes.Client)
	if err != nil {
		tx.Rollback()
		ur.log.Debug("Не удалось создать психолога по причине: " + err.Error())
		return nil, &models.ErrorResponse{ErrorMessage: "Адрес электронной почты занят", ErrorCode: 400}
	}

	return user, nil
}

func (ur *UserRepo) ContinueSignUp(csu *models.ContinueSignUp) *models.ErrorResponse {

	_, err := ur.db.Exec(`
	UPDATE Users
		SET 
		    avatarurl=$1, description=$2
	WHERE
		id = $3
`, csu.Avatar, csu.Description, csu.UserID)
	if err != nil {
		ur.log.Debug("Не удалось продолжить регистрацию: " + err.Error())
		return &models.ErrorResponse{ErrorMessage: "Не удалось продолжить регистрацию", ErrorCode: 500}
	}

	return nil
}

func (ur *UserRepo) UpdateProfile(userOrig, userUpd *models.User) *models.ErrorResponse {
	tx, err := ur.db.Begin()
	if err != nil {
		ur.log.Error(err.Error())
		return &models.ErrorResponse{ErrorMessage: "Не удалось подготовить базу", ErrorCode: 500}
	}
	defer tx.Commit()

	if userOrig.Firstname != userUpd.Firstname {
		_, err = tx.Exec(`

		UPDATE Users
		    SET firstname=$1
		Where 
		    id =$2
`, userUpd.Firstname, userUpd.ID)
		if err != nil {
			tx.Rollback()
			ur.log.Debug("не удалось обновить имя: " + err.Error())
			return &models.ErrorResponse{ErrorMessage: "не удалось обновить имя", ErrorCode: 400}
		}
	}

	if userOrig.Lastname != userUpd.Lastname {
		_, err = tx.Exec(`
		UPDATE Users
		    SET lastname=$1
		Where 
		    id =$2
`, userUpd.Lastname, userUpd.ID)
		if err != nil {
			tx.Rollback()
			ur.log.Debug("не удалось обновить фамилию: " + err.Error())
			return &models.ErrorResponse{ErrorMessage: "не удалось обновить фамилию", ErrorCode: 400}
		}
	}

	if userOrig.Username != userUpd.Username {
		_, err = tx.Exec(`
		UPDATE Users
		    SET username=$1
		Where 
		    id =$2
`, userUpd.Username, userUpd.ID)
		if err != nil {
			tx.Rollback()
			ur.log.Debug("не удалось обновить ник: " + err.Error())
			return &models.ErrorResponse{ErrorMessage: "ник уже занят, пожалуйста используйте другой", ErrorCode: 400}
		}
	}
	if userOrig.Description != userUpd.Description {
		_, err = tx.Exec(`
		UPDATE Users
		    SET description=$1
		Where 
		    id =$2
`, userUpd.Description, userUpd.ID)
		if err != nil {
			tx.Rollback()
			ur.log.Debug("не удалось обновить описание профиля: " + err.Error())
			return &models.ErrorResponse{ErrorMessage: "не удалось обновить описание профиля", ErrorCode: 400}
		}
	}

	if userOrig.Age != userUpd.Age {
		_, err = tx.Exec(`
		UPDATE Users
		    SET age=$1
		Where 
		    id =$2
`, userUpd.Age, userUpd.ID)
		if err != nil {
			tx.Rollback()
			ur.log.Debug("не удалось обновить возраст: " + err.Error())
			return &models.ErrorResponse{ErrorMessage: "не удалось обновить возраст", ErrorCode: 400}
		}
	}
	return nil
}
