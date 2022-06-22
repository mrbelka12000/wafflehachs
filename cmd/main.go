package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wafflehacks/database"
	"wafflehacks/internal/app"
	"wafflehacks/tools/storage"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		println(err.Error())
		return
	}
	log := logger.Sugar()
	defer log.Sync()

	db, err := database.GetConnection()
	if err != nil {
		log.Debug(err.Error())
		return
	}

	go database.DeleteExpiredCookie(db, log)

	srv := app.Initialize(db, log)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			done <- os.Interrupt
			log.Info(err.Error())
			return
		}
	}()

	log.Info("Server started on port: " + os.Getenv("PORT"))

	<-done
	log.Info("Server Stopped")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling
		if err = db.Close(); err != nil {
			log.Debug(err.Error())
		} else {
			log.Info("Connection to database successfully closed")
		}

		if err = os.Remove(storage.GoogleConfigFileName); err != nil {
			log.Debug(err.Error())
		} else {
			log.Info("Temp files removed")
		}

		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Debugf("Server Shutdown Failed:%+v", err)
		return
	}

	log.Info("Server Exited Properly")
}
