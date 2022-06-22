package http

import (
	"encoding/json"
	"net/http"
	"wafflehacks/entities/response"

	request "wafflehacks/entities/requests"
	"wafflehacks/models"
	"wafflehacks/tools"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := request.ClientSignInRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.SendErrorResponse(w, "serailization failed: "+err.Error(), 400)
		return
	}

	u := req.Build()

	user, resp := h.srv.User.CanLogin(u)
	if resp != nil {
		response.SendErrorResponse(w, "not found", 400)
		return
	}

	s := &models.SessionResponse{
		ID:     user.ID,
		Cookie: tools.GetRandomString(),
	}

	if err := h.srv.CreateSession(s); err != nil {
		h.log.Debug("create session failed: " + err.ErrorMessage)
		response.SendErrorResponse(w, err.ErrorMessage, err.ErrorCode)
		return
	}
	w.Write([]byte(tools.MakeJsonString(s)))
}
