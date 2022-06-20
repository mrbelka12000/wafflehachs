package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"wafflehacks/tools"
)

func (h *Handler) GetPsycho(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		SendErrorResponse(w, "not found", 404)
		return
	}

	psycho, resp := h.srv.GetByUsername(username)
	if resp != nil {
		SendErrorResponse(w, resp.ErrorMessage, resp.ErrorCode)
		return
	}
	w.Write([]byte(tools.MakeJsonString(psycho)))
}
