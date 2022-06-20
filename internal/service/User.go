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

func (us *UserService) CanLogin(user *models.User) (*models.User, *models.ErrorResponse) {
	u, resp := us.repo.CanLogin(user)
	if resp != nil {
		us.log.Error("user not found")
		return nil, resp
	}

	err := bcrypt.CompareHashAndPassword([]byte(*u.Password), []byte(*user.Password+os.Getenv("Salt")))
	if err == nil {
		return u, nil
	}
	us.log.Error("wrong password")
	return nil, &models.ErrorResponse{ErrorMessage: "wrong password", ErrorCode: 400}
}

func (us *UserService) SignUp(user *models.User) (*models.User, *models.ErrorResponse) {
	if err := validate.ValidatingPassword(*user.Password); err != nil {
		us.log.Debug(err.Error())
		return nil, &models.ErrorResponse{ErrorMessage: err.Error(), ErrorCode: 400}
	}

	password, err := bcrypt.GenerateFromPassword([]byte(*user.Password+os.Getenv("Salt")), bcrypt.DefaultCost)
	if err != nil {
		us.log.Debug("Ошибка шифрования пароля")
		return nil, &models.ErrorResponse{ErrorMessage: "Не уда3лось зафишровать данные", ErrorCode: 500}
	}
	user.Password = tools.GetPointerString(string(password))

	if user.Age < 16 {
		return nil, &models.ErrorResponse{ErrorMessage: "Не подходите по возрасту", ErrorCode: 400}
	}

	if err = validate.ValidatingEmail(user.Email); err != nil {
		us.log.Debug(fmt.Sprintf("Адрес %v не подходит по требованиям ", user.Email))
		return nil, &models.ErrorResponse{ErrorMessage: err.Error(), ErrorCode: 400}
	}
	return us.repo.User.SignUp(user)
}

func (us *UserService) ContinueSignUp(csu *models.ContinueSignUp) *models.ErrorResponse {
	return us.repo.ContinueSignUp(csu)
}
