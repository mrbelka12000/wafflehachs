package repository

import (
	"database/sql"

	m "wafflehacks/models"

	"go.uber.org/zap"
)

type Psychologist interface {
	SignUp(psychoId int) *m.ErrorResponse
	GetAll() ([]m.Psychologist, *m.ErrorResponse)
	GetByUsername(username string) (*m.Psychologist, *m.ErrorResponse)
}

type Client interface {
	SignUp(clientId int) *m.ErrorResponse
}

type User interface {
	CanLogin(user *m.User) (*m.User, *m.ErrorResponse)
	SignUp(user *m.User) (*m.User, *m.ErrorResponse)
	ContinueSignUp(csu *m.ContinueSignUp) *m.ErrorResponse
}

type Session interface {
	CreateSession(session *m.SessionResponse) *m.ErrorResponse
}

type Repository struct {
	Psychologist
	Client
	User
	Session
}

func NewRepo(db *sql.DB, log *zap.SugaredLogger) *Repository {
	return &Repository{
		Psychologist: newPsycho(db, log),
		Client:       newClient(db, log),
		User:         newUser(db, log),
		Session:      newSession(db, log),
	}
}
