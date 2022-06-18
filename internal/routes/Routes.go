package routes

import (
	"net/http"
	"wafflehacks/internal/handler"

	"github.com/gorilla/mux"
)

func SetUpMux(h *handler.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/signup", h.SignUp).Methods(http.MethodPost)
	r.HandleFunc("/api/signin", h.SignIn).Methods(http.MethodPost)

	r.HandleFunc("/api/main", h.MainPage).Methods(http.MethodPost)
	return r
}
