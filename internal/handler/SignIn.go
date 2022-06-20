package handler

import (
	"encoding/json"
	"net/http"
	request "wafflehacks/entities/requests"
	"wafflehacks/tools"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := request.ClientSignInRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendErrorResponse(w, "serailization failed: "+err.Error(), 400)
		return
	}

	u := req.Build()

	user, resp := h.srv.User.CanLogin(u)
	if resp != nil {
		SendErrorResponse(w, "not found", 400)
		return
	}
	w.Write([]byte(tools.MakeJsonString(user)))
}
