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
`, csu.Avatar, csu.Description, csu.ID)

	if err != nil {
		ur.log.Debug("Не удалось продолжить регистрацию: " + err.Error())
		return &models.ErrorResponse{ErrorMessage: "Не удалось продолжить регистрацию", ErrorCode: 500}
	}

	return nil
}
