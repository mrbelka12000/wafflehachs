package handler

import (
	"net/http"
	"wafflehacks/models"
	"wafflehacks/tools"
)

func SendErrorResponse(w http.ResponseWriter, errorMsg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	if code == 0 {
		code = http.StatusInternalServerError
	}
	w.WriteHeader(code)

	er := &models.ErrorResponse{ErrorMessage: errorMsg, ErrorCode: code}
	w.Write([]byte(tools.MakeJsonString(er)))
}
