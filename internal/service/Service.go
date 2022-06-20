package service

import (
	"wafflehacks/internal/repository"
	m "wafflehacks/models"

	"go.uber.org/zap"
)

type Psychologist interface {
	SignUp(psychoId int) *m.ErrorResponse
	GetAll() ([]m.Psychologist, *m.ErrorResponse)
	GetByUsername(username string) (*m.Psychologist, *m.ErrorResponse)
	UpdateBusyMode(mode string, psychoId int) *m.ErrorResponse
}

type Client interface {
	SignUp(clientId int) *m.ErrorResponse
}

type User interface {
	CanLogin(email *m.User) (*m.User, *m.ErrorResponse)
	SignUp(user *m.User, usertype string) (*m.User, *m.ErrorResponse)
	ContinueSignUp(csu *m.ContinueSignUp) *m.ErrorResponse
	UpdateProfile(userOrig, userUpd *m.User) *m.ErrorResponse
}

type Session interface {
	CreateSession(session *m.SessionResponse) *m.ErrorResponse
	GetUserByCookie(cookie string) (*m.User, *m.ErrorResponse)
}

type Service struct {
	Psychologist
	Client
	User
	Session
}

func NewService(repo *repository.Repository, log *zap.SugaredLogger) *Service {
	return &Service{
		Psychologist: newPsycho(repo, log),
		Client:       newClient(repo, log),
		User:         newUser(repo, log),
		Session:      newSession(repo, log),
	}
}
