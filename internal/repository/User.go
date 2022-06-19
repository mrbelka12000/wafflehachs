package repository

import (
	"database/sql"
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

func (u *UserRepo) CanLogin(user *models.User) (*models.User, *models.ErrorResponse) {
	User := &models.User{}

	err := u.db.QueryRow(`
		SELECT 
			id, Password
		FROM
			users
		WHERE
			Email = $1	
		`, user.Email).Scan(&User.ID, &User.Password)
	if err != nil {
		u.log.Debug("email not found : " + err.Error())
		return nil, &models.ErrorResponse{ErrorMessage: "getUser failed", ErrorCode: 500}
	}

	return User, nil
}
