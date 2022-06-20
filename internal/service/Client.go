package service

import (
	"go.uber.org/zap"
	"wafflehacks/internal/repository"
	"wafflehacks/models"
)

type ClientService struct {
	repo *repository.Repository
	log  *zap.SugaredLogger
}

func newClient(repo *repository.Repository, log *zap.SugaredLogger) *ClientService {
	return &ClientService{repo, log}
}

func (cs *ClientService) SignUp(clientId int) *models.ErrorResponse {

	return cs.repo.Client.SignUp(clientId)
}
