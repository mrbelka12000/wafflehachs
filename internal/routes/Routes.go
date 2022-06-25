package routes

import (
	"net/http"
	h "wafflehacks/internal/handler/http"
	ws "wafflehacks/internal/handler/websocket"

	"github.com/gorilla/mux"
)

func SetUpMux(handler *h.Handler, websocket *ws.Handler) *mux.Router {
	r := mux.NewRouter()

	//Http Handlers
	r.HandleFunc("/api/signup", handler.SignUp).Methods(http.MethodPost)
	r.HandleFunc("/api/signin", handler.SignIn).Methods(http.MethodPost)

	r.HandleFunc("/api/main", handler.MainPage).Methods(http.MethodGet)
	r.HandleFunc("/api/upload", handler.Upload).Methods(http.MethodPost)

	r.HandleFunc("/api/user", handler.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/api/psychologist/{username}", handler.GetPsycho).Methods(http.MethodGet)
	r.HandleFunc("/api/psychologist/change/{busymode}", handler.UpdateBusyMode).Methods(http.MethodPost)

	r.HandleFunc("/api/room/register", nil)

	//WebSocket Handlers
	r.HandleFunc("/api/room/{id}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "main.html")
	})

	r.HandleFunc("/api/ws/room/{id}", websocket.GetConnection).Methods(http.MethodGet)
	return r
}
