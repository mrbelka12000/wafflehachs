package handler

import (
	"encoding/json"
	"net/http"
	request "wafflehacks/entities/requests"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	subject := r.FormValue("subject")
	switch subject {
	case "client":

		clientSignUpReq := request.ClientSignUpRequest{}
		if err := json.NewDecoder(r.Body).Decode(&clientSignUpReq); err != nil {
			SendErrorResponse(w, "Ошибка дессириализации: "+err.Error(), 400)
			h.log.Debug("Ошибка дессириализации: " + err.Error())
			return
		}

		if err := clientSignUpReq.Validate(); err != nil {
			SendErrorResponse(w, err.Error(), 400)
			h.log.Debug(err.Error())
			return
		}

		client := clientSignUpReq.Build()
		client, resp := h.srv.Client.SignUp(client)
		if resp != nil {
			SendErrorResponse(w, resp.ErrorMessage, resp.ErrorCode)
			h.log.Debug(resp.ErrorMessage)
			return
		}

	case "psychologist":

		psychoSignUpReq := request.PsychoSignUpRequest{}
		if err := json.NewDecoder(r.Body).Decode(&psychoSignUpReq); err != nil {
			SendErrorResponse(w, "Ошибка дессириализации: "+err.Error(), 400)
			h.log.Debug("Ошибка дессириализации: " + err.Error())
			return
		}

		if err := psychoSignUpReq.Validate(); err != nil {
			SendErrorResponse(w, err.Error(), 400)
			h.log.Debug(err.Error())
			return
		}

		psycho := psychoSignUpReq.Build()

		psycho, resp := h.srv.Psychologist.SignUp(psycho)
		if resp != nil {
			SendErrorResponse(w, resp.ErrorMessage, resp.ErrorCode)
			h.log.Debug(resp.ErrorMessage)
			return
		}

	default:
		SendErrorResponse(w, "Неизвестный тип регистрации", 400)
		h.log.Debug("Неизвестный тип регистрации")
		return
	}
}
