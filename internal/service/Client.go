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

type ClientService struct {
	repo *repository.Repository
	log  *zap.SugaredLogger
}

func newClient(repo *repository.Repository, log *zap.SugaredLogger) *ClientService {
	return &ClientService{repo, log}
}

func (cs *ClientService) SignUp(client *models.Client) (*models.Client, *models.ErrorResponse) {
	if err := validate.ValidatingPassword(*client.Password); err != nil {
		cs.log.Debug(err.Error())
		return nil, &models.ErrorResponse{ErrorMessage: err.Error()}
	}

	if err := validate.ValidatingEmail(client.Email); err != nil {
		cs.log.Debug(fmt.Sprintf("Адрес %v не подходит по требованиям ", client.Email))
		return nil, &models.ErrorResponse{ErrorMessage: err.Error(), ErrorCode: 400}
	}
	password, err := bcrypt.GenerateFromPassword([]byte(*client.Password+os.Getenv("Salt")), bcrypt.DefaultCost)
	if err != nil {
		cs.log.Debug("Ошибка шифрования пароля")
		return nil, &models.ErrorResponse{ErrorMessage: "Не уда3лось зафишровать данные", ErrorCode: 500}
	}
	client.Password = tools.GetPointerString(string(password))

	if client.Age < 16 {
		return nil, &models.ErrorResponse{ErrorMessage: "Не подходите по возрасту", ErrorCode: 400}
	}

	return cs.repo.Client.SignUp(client)
}
