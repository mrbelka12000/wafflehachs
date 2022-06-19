package repository

import (
	"database/sql"
	m "wafflehacks/models"

	"go.uber.org/zap"
)

type Psychologist interface {
	SignUp(psycho *m.Psychologist) (*m.Psychologist, *m.ErrorResponse)
	GetAll() ([]m.Psychologist, *m.ErrorResponse)
}

type Client interface {
	SignUp(client *m.Client) (*m.Client, *m.ErrorResponse)
}

type User interface {
	CanLogin(user *m.User) (*m.User, *m.ErrorResponse)
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
