package handler

import (
	"encoding/json"
	"net/http"

	request "wafflehacks/entities/requests"
	"wafflehacks/models"
	"wafflehacks/tools"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	req := request.ClientSignInRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendErrorResponse(w, "serailization failed: "+err.Error(), 400)
		h.log.Debug("serailization failed: "+err.Error(), 400)
		return
	}

	u := req.Build()

	user, resp := h.srv.User.GetUser(u)
	if resp != nil {
		h.log.Debug(resp)
		SendErrorResponse(w, "not found", 400)
		return
	}

	s := &models.SessionResponse{
		ID:     user.ID,
		Cookie: tools.GetRandomString(),
	}

	if err := h.srv.CreateSession(s); err != nil {
		h.log.Debug("create session failed: " + err.ErrorMessage)
		SendErrorResponse(w, err.ErrorMessage, err.ErrorCode)
		return
	}
}
