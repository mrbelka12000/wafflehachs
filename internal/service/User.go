package service

import (
	"os"
	"wafflehacks/internal/repository"
	"wafflehacks/models"

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
		us.log.Debug("user not found")
		return nil, resp
	}

	err := bcrypt.CompareHashAndPassword([]byte(*u.Password), []byte(*user.Password+os.Getenv("Salt")))
	if err == nil {
		return u, nil
	}
	us.log.Debug("wrong password")
	return nil, &models.ErrorResponse{ErrorMessage: "wrong password", ErrorCode: 400}
}
