package server

import (
	"net/http"
	"os"
	"time"
	"wafflehacks/internal/handler"
	"wafflehacks/internal/routes"
)

func NewServer(h *handler.Handler) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &http.Server{
		WriteTimeout: time.Duration(25 * time.Second),
		ReadTimeout:  time.Duration(25 * time.Second),
		Handler:      routes.SetUpMux(h),
		Addr:         ":" + port,
	}
}
