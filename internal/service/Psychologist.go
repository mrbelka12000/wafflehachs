package service

import (
	"go.uber.org/zap"
	"wafflehacks/internal/repository"
	"wafflehacks/models"
)

type PsychoService struct {
	repo *repository.Repository
	log  *zap.SugaredLogger
}

func newPsycho(repo *repository.Repository, log *zap.SugaredLogger) *PsychoService {
	return &PsychoService{repo, log}
}

func (ps *PsychoService) SignUp(psychoId int) *models.ErrorResponse {
	return ps.repo.Psychologist.SignUp(psychoId)
}
func (ps *PsychoService) GetAll() ([]models.Psychologist, *models.ErrorResponse) {
	return ps.repo.GetAll()
}

func (ps *PsychoService) GetByUsername(username string) (*models.Psychologist, *models.ErrorResponse) {
	return ps.repo.GetByUsername(username)
}
