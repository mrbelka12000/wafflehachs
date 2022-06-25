package service

import (
	"go.uber.org/zap"
	"wafflehacks/internal/repository"
	"wafflehacks/models"
)

type RoomService struct {
	repo *repository.Repository
	log  *zap.SugaredLogger
}

func newRoom(repo *repository.Repository, log *zap.SugaredLogger) *RoomService {
	return &RoomService{repo, log}
}

func (rs *RoomService) CreateRoom(room *models.Room) *models.ErrorResponse {
	return rs.repo.CreateRoom(room)
}
