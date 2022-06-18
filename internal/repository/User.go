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
