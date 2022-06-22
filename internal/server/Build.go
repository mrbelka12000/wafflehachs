package server

import (
	"net/http"
	"os"
	"time"
	h "wafflehacks/internal/handler/http"
	ws "wafflehacks/internal/handler/websocket"
	"wafflehacks/internal/routes"
)

func NewServer(h *h.Handler, ws *ws.Handler) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &http.Server{
		WriteTimeout: 25 * time.Second,
		ReadTimeout:  25 * time.Second,
		IdleTimeout:  25 * time.Second,
		Handler:      routes.SetUpMux(h, ws),
		Addr:         ":" + port,
	}
}
