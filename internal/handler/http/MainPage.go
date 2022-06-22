package http

import (
	"net/http"
	"wafflehacks/entities/response"
	"wafflehacks/tools"
)

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	psychos, resp := h.srv.GetAll()
	if resp != nil {
		response.SendErrorResponse(w, resp.ErrorMessage, resp.ErrorCode)
		return
	}

	w.Write([]byte(tools.MakeJsonString(psychos)))
}
