package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wafflehacks/database"
	"wafflehacks/internal/handler"
	"wafflehacks/internal/repository"
	"wafflehacks/internal/server"
	"wafflehacks/internal/service"
	"wafflehacks/tools"
	"wafflehacks/tools/storage"

	"go.uber.org/zap"
)

func Initialize(db *sql.DB, log *zap.SugaredLogger) *http.Server {
	repo := repository.NewRepo(db, log)
	serv := service.NewService(repo, log)
	handler := handler.NewHandler(serv, log)
	return server.NewServer(handler)
}

func Run(log *zap.SugaredLogger) {

	db, err := database.GetConnection()
	if err != nil {
		log.Debug(err.Error())
		return
	}

	go database.DeleteExpiredCookie(db, log)

	srv := Initialize(db, log)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			done <- syscall.SIGQUIT
			log.Debug(err.Error())
			return
		}
	}()

	log.Info("Server started on port: " + os.Getenv("PORT"))

	signalFromSystem := <-done
	log.Info("Server Stopped")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		if err = db.Close(); err != nil {
			log.Debug(err.Error())
		} else {
			log.Info("Connection to database successfully closed")
		}

		if tools.CheckSignal(signalFromSystem) {
			if err = os.Remove(storage.GoogleConfigFileName); err != nil {
				log.Debug(err.Error())
			} else {
				log.Info("Temp files removed")
			}
		} else {
			log.Debugf("Unknown signal called: %v", signalFromSystem)
		}

		cancel()
	}()

	if err = srv.Shutdown(ctx); err != nil {
		log.Debugf("Server Shutdown Failed:%+v", err)
		return
	}
}

func init() {
	tools.Loadenv()
	res := fmt.Sprintf(`
{
	"type": "%v",
	"project_id": "%v",
	"private_key_id": "%v",
	"private_key": "%v",
	"client_email": "%v",
	"client_id": "%v",
	"auth_uri": "%v",
	"token_uri": "%v",
	"auth_provider_x509_cert_url": "%v",
	"client_x509_cert_url": "%v"
}
`, os.Getenv("type"), os.Getenv("project_id"), os.Getenv("private_key_id"),
		os.Getenv("private_key"), os.Getenv("client_email"), os.Getenv("client_id"),
		os.Getenv("auth_uri"), os.Getenv("token_uri"), os.Getenv("auth_provider_x509_cert_url"),
		os.Getenv("client_x509_cert_url"))
	file, err := os.Create(storage.GoogleConfigFileName)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write([]byte(res))
	if err != nil {
		log.Fatal(err)
	}
}
