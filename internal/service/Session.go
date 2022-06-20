package service

import (
	"wafflehacks/internal/repository"
	"wafflehacks/models"

	"go.uber.org/zap"
)

type SessionService struct {
	repo *repository.Repository
	log  *zap.SugaredLogger
}

func newSession(repo *repository.Repository, log *zap.SugaredLogger) *SessionService {
	return &SessionService{
		repo: repo,
		log:  log,
	}
}

func (s *SessionService) CreateSession(session *models.SessionResponse) *models.ErrorResponse {
	return s.repo.CreateSession(session)
}

func (s *SessionService) GetUserIdByCookie(cookie string) (int, *models.ErrorResponse) {
	return s.repo.GetUserIdByCookie(cookie)
}
