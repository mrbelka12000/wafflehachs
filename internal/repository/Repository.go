package repository

import (
	"database/sql"
	"wafflehacks/models"
	m "wafflehacks/models"

	"go.uber.org/zap"
)

type Psychologist interface {
	SignUp(psycho *m.Psychologist) (*m.Psychologist, *models.ErrorResponse)
}

type Client interface {
	SignUp(client *m.Client) (*m.Client, *models.ErrorResponse)
}

type User interface {
	GetType(email string) (string, *models.ErrorResponse)
}

type Repository struct {
	Psychologist
	Client
	User
}

func NewRepo(db *sql.DB, log *zap.SugaredLogger) *Repository {
	return &Repository{
		Psychologist: newPsycho(db, log),
		Client:       newClient(db, log),
		User:         newUser(db, log),
	}
}
