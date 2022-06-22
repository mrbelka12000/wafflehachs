package http

import (
	"fmt"
	"net/http"
	request "wafflehacks/entities/requests"
	"wafflehacks/entities/response"
	"wafflehacks/tools"
	"wafflehacks/tools/storage"
)

const MaxSize = 10 << 20

func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	user, err := h.getUserId(r)
	if err != nil {
		response.SendErrorResponse(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	r.ParseMultipartForm(10 << 20)
	userConfirm := &request.UserSignUpContinueRequest{}
	var haveError bool
	file, handler, err := r.FormFile("avatar")
	if err != nil {
		haveError = true
	}
	if !haveError {
		if handler.Size >= MaxSize {
			response.SendErrorResponse(w, "Изображение весит больше чем положено, 20mb", 400)
			h.log.Debug("Слишком большой файл")
			return
		}
		defer file.Close()

		if fileType, ok := tools.IsValidType(file); !ok {
			response.SendErrorResponse(w, fmt.Sprintf("Разрешение %v не поддерживается", fileType), 400)
			h.log.Debug(fmt.Sprintf("Разрешение %v не поддерживается", fileType))
			return
		}

		file.Seek(0, 0)

		filename := tools.GetRandomString()
		if err := storage.UploadFile(file, filename); err != nil {
			response.SendErrorResponse(w, "Не удалось загрузить файл ", 500)
			h.log.Debug("Не удалось загрузить файл по причине: " + err.Error())
			return
		}

		userConfirm.Avatar = filename
	}
	desc := r.FormValue("description")
	userConfirm.Description = desc

	if userConfirm.Handle() {
		response.SendErrorResponse(w, "Пустые данные, добавлять ничего не будет", http.StatusOK)
		return
	}

	confirmSignUp := userConfirm.Build(user.ID)
	resp := h.srv.ContinueSignUp(confirmSignUp)
	if resp != nil {
		response.SendErrorResponse(w, resp.ErrorMessage, resp.ErrorCode)
		return
	}

	response.SendErrorResponse(w, "", http.StatusOK)
}
