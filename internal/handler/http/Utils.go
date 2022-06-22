package http

import (
	"errors"
	"net/http"
	"wafflehacks/models"
)

var unknownUser = errors.New("нe известный пользователь")

func (h *Handler) getUserId(r *http.Request) (*models.User, error) {
	cookie := r.Header.Get("session")
	if cookie == "" {
		h.log.Debug("Не известный пользователь: ", r.RemoteAddr)
		return nil, unknownUser
	}
	user, resp := h.srv.GetUserByCookie(cookie)
	if resp != nil {
		h.log.Debug(resp.ErrorMessage)
		return nil, unknownUser
	}
	return user, nil
}
