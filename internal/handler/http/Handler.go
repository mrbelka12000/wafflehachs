package http

import (
	"wafflehacks/internal/service"

	"go.uber.org/zap"
)

type Handler struct {
	srv *service.Service
	log *zap.SugaredLogger
}

func NewHandler(srv *service.Service, log *zap.SugaredLogger) *Handler {
	return &Handler{srv, log}
}
