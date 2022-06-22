package websocket

import (
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"wafflehacks/internal/service"
)

type Handler struct {
	srv      *service.Service
	log      *zap.SugaredLogger
	Upgrader websocket.Upgrader
	Hub      *Hub
}

func NewHandler(srv *service.Service, log *zap.SugaredLogger) *Handler {
	return &Handler{
		srv: srv,
		log: log,
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		Hub: NewHub(),
	}
}
