package http

import (
	"encoding/json"
	"net/http"
	request "wafflehacks/entities/requests"
	"wafflehacks/entities/response"
	"wafflehacks/entities/usertypes"
	"wafflehacks/models"
	"wafflehacks/tools"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userSignUpReq := &request.UserSignUpRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userSignUpReq); err != nil {
		response.SendErrorResponse(w, "Ошибка дессириализации: "+err.Error(), 400)
		return
	}
	user := userSignUpReq.Build()

	subject := r.FormValue("subject")
	if subject != usertypes.Client && subject != usertypes.Psycho {
		response.SendErrorResponse(w, "Неизвестный тип регистрации", 400)
		return
	}

	user, resp := h.srv.User.SignUp(user, subject)
	if resp != nil {
		response.SendErrorResponse(w, resp.ErrorMessage, resp.ErrorCode)
		return
	}
	switch subject {
	case "client":
		resp = h.srv.Client.SignUp(user.ID)
	case "psychologist":
		resp = h.srv.Psychologist.SignUp(user.ID)
	}

	if resp != nil {
		response.SendErrorResponse(w, resp.ErrorMessage, resp.ErrorCode)
		return
	}
	s := &models.SessionResponse{
		ID:     user.ID,
		Cookie: tools.GetRandomString(),
	}
	resp = h.srv.CreateSession(s)
	if resp != nil {
		response.SendErrorResponse(w, resp.ErrorMessage, resp.ErrorCode)
		return
	}
	w.Write([]byte(tools.MakeJsonString(s)))
}
