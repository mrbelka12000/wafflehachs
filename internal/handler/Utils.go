package handler

import (
	"errors"
	"net/http"
)

var unknownUser = errors.New("нe известный пользователь")

func (h *Handler) getUserId(r *http.Request) (int, error) {
	cookie := r.Header.Get("session")
	if cookie == "" {
		h.log.Debug("Не известный пользователь: ", r.RemoteAddr)
		return 0, unknownUser
	}
	user, resp := h.srv.GetUserByCookie(cookie)
	if resp != nil {
		h.log.Debug(resp.ErrorMessage)
		return 0, unknownUser
	}
	return user.ID, nil
}
