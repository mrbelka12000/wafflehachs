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

func (u *UserRepo) GetType(email string) (string, *models.ErrorResponse) {
	var role string

	err := u.db.QueryRow(`
	SELECT 
		Role 
	FROM 
		UserType 
	WHERE 
		Email = $1
	`, email).Scan(&role)
	if err != nil {
		u.log.Debug("gmail не найден")
		return "", &models.ErrorResponse{ErrorMessage: "gmail не найден", ErrorCode: 400}
	}

	return role, nil
}

func (u *UserRepo) GetUser(user *models.User) (*models.User, *models.ErrorResponse) {
	User := &models.User{}

	role, resp := u.GetType(user.Email)
	if resp != nil {
		u.log.Debug("gettype failed")
		return nil, &models.ErrorResponse{ErrorMessage: "gettype failed", ErrorCode: 500}
	}

	switch role {
	case usertypes.Client:
		err := u.db.QueryRow(`
		SELECT 
			id, Password
		FROM
			Clients
		WHERE
			Email = $1	
		`, user.Email).Scan(&User.ID, &User.Password)
		if err != nil {
			u.log.Debug("email not found : " + err.Error())
			return nil, &models.ErrorResponse{ErrorMessage: "getUser failed", ErrorCode: 500}
		}
	case usertypes.Psycho:
		err := u.db.QueryRow(`
		SELECT 
			id,Password
		FROM
			Psychologists
		WHERE
			Email = $1	
		`, user.Email).Scan(&User.ID, &User.Password)
		if err != nil {
			u.log.Debug("email not found: " + err.Error())
			return nil, &models.ErrorResponse{ErrorMessage: "getUser failed", ErrorCode: 500}
		}
	}

	return User, nil
}
