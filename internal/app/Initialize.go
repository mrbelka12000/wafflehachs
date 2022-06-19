package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"wafflehacks/entities/storage"
	"wafflehacks/internal/handler"
	"wafflehacks/internal/repository"
	"wafflehacks/internal/server"
	"wafflehacks/internal/service"
	"wafflehacks/tools"

	"go.uber.org/zap"
)

func Initialize(db *sql.DB, log *zap.SugaredLogger) *http.Server {
	repo := repository.NewRepo(db, log)
	serv := service.NewService(repo, log)
	handler := handler.NewHandler(serv, log)
	return server.NewServer(handler)
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
