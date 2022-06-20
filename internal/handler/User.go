package handler

import (
	"net/http"
	request "wafflehacks/entities/requests"
)

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := h.getUserId(r)
	if err != nil {
		SendErrorResponse(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	uu := &request.UpdateUserRequest{}

	uu, err = uu.BuildRequest(id, r)
	if err != nil {
		SendErrorResponse(w, err.Error(), 400)
		h.log.Debug(err.Error())
		return
	}

	user := uu.Build()

	resp := h.srv.UpdateProfile(user, r.Header.Get("session"))
	if resp != nil {
		SendErrorResponse(w, resp.ErrorMessage, resp.ErrorCode)
		h.log.Debug(resp.ErrorMessage)
		return
	}

	SendErrorResponse(w, "", 200)
}
