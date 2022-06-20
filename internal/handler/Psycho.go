package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	request "wafflehacks/entities/requests"
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

func (h *Handler) UpdateBusyMode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, err := h.getUserId(r)
	if err != nil {
		SendErrorResponse(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	bm := request.BusyMode(vars["busymode"])
	if err = bm.CheckForExists(); err != nil {
		SendErrorResponse(w, err.Error(), 400)
		h.log.Debug(err.Error())
		return
	}

	resp := h.srv.UpdateBusyMode(string(bm), user.ID)
	if resp != nil {
		SendErrorResponse(w, resp.ErrorMessage, resp.ErrorCode)
		h.log.Debug("не удалось обновить данные по причине: " + resp.ErrorMessage)
		return
	}

	SendErrorResponse(w, "", http.StatusOK)
}
