package service

import (
	"fmt"
	"os"
	"wafflehacks/internal/repository"
	"wafflehacks/internal/service/validate"
	"wafflehacks/models"
	"wafflehacks/tools"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type PsychoService struct {
	repo *repository.Repository
	log  *zap.SugaredLogger
}

func newPsycho(repo *repository.Repository, log *zap.SugaredLogger) *PsychoService {
	return &PsychoService{repo, log}
}

func (ps *PsychoService) SignUp(psycho *models.Psychologist) (*models.Psychologist, *models.ErrorResponse) {
	if err := validate.ValidatingPassword(*psycho.Password); err != nil {
		ps.log.Debug(err.Error())
		return nil, &models.ErrorResponse{ErrorMessage: err.Error(), ErrorCode: 400}
	}

	password, err := bcrypt.GenerateFromPassword([]byte(*psycho.Password+os.Getenv("Salt")), bcrypt.DefaultCost)
	if err != nil {
		ps.log.Debug("Ошибка шифрования пароля")
		return nil, &models.ErrorResponse{ErrorMessage: "Не уда3лось зафишровать данные", ErrorCode: 500}
	}
	psycho.Password = tools.GetPointerString(string(password))

	if psycho.Age < 16 {
		return nil, &models.ErrorResponse{ErrorMessage: "Не подходите по возрасту", ErrorCode: 400}
	}

	if err = validate.ValidatingEmail(psycho.Email); err != nil {
		ps.log.Debug(fmt.Sprintf("Адрес %v не подходит по требованиям ", psycho.Email))
		return nil, &models.ErrorResponse{ErrorMessage: err.Error(), ErrorCode: 400}
	}
	return ps.repo.Psychologist.SignUp(psycho)
}

func (ps *PsychoService) GetAll() ([]models.Psychologist, *models.ErrorResponse) {
	return ps.repo.GetAll()
}
