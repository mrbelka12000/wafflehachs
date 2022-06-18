package service

import (
	"wafflehacks/internal/repository"
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

type Service struct {
	Psychologist
	Client
	User
}

func NewService(repo *repository.Repository, log *zap.SugaredLogger) *Service {
	return &Service{
		Psychologist: newPsycho(repo, log),
		Client:       newClient(repo, log),
		User:         newUser(repo, log),
	}
}
