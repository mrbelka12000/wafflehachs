package service

import (
	"wafflehacks/internal/repository"
	"wafflehacks/models"

	"go.uber.org/zap"
)

type UserService struct {
	repo *repository.Repository
	log  *zap.SugaredLogger
}

func newUser(repo *repository.Repository, log *zap.SugaredLogger) *UserService {
	return &UserService{
		repo: repo,
		log:  log,
	}
}

func (us *UserService) GetType(email string) (string, *models.ErrorResponse) {
	return us.repo.GetType(email)
}
