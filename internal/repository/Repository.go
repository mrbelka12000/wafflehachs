package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type Psychologyst interface {
	SignUp()
}

type Client interface {
	SignUp()
}

type Repository struct {
	Psychologyst
	Client
}

func NewRepo(db *sql.DB, log *zap.SugaredLogger) *Repository {
	return &Repository{
		Psychologyst: NewPsycho(db, log),
		Client:       NewClient(db, log),
	}
}
