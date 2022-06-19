package handler

import (
	"encoding/json"
	"net/http"
	request "wafflehacks/entities/requests"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userSignUpReq := &request.UserSignUpRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userSignUpReq); err != nil {
		SendErrorResponse(w, "Ошибка дессириализации: "+err.Error(), 400)
		h.log.Debug("Ошибка дессириализации: " + err.Error())
		return
	}
	user := userSignUpReq.Build()

	subject := r.FormValue("subject")
	if subject != "client" && subject != "psychologist" {
		SendErrorResponse(w, "Неизвестный тип регистрации", 400)
		h.log.Debug("Неизвестный тип регистрации")
		return
	}

	user, resp := h.srv.User.SignUp(user)
	if resp != nil {
		SendErrorResponse(w, resp.ErrorMessage, resp.ErrorCode)
		h.log.Debug(resp.ErrorMessage)
		return
	}
	switch subject {
	case "client":
		resp = h.srv.Client.SignUp(user.ID)
	case "psychologist":
		resp = h.srv.Psychologist.SignUp(user.ID)
	}

	if resp != nil {
		SendErrorResponse(w, resp.ErrorMessage, resp.ErrorCode)
		h.log.Debug(resp.ErrorMessage)
		return
	}

	SendErrorResponse(w, "", http.StatusOK)
}
