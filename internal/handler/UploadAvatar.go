package handler

import (
	"fmt"
	"net/http"
	request "wafflehacks/entities/requests"
	"wafflehacks/entities/storage"
	"wafflehacks/tools"
)

const MaxSize = 10 << 20

func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {

	//cookie := r.Header.Get("sessionCookie")
	r.ParseMultipartForm(10 << 20)
	userConfirm := &request.UserSignUpContinueRequest{}
	var haveError bool
	file, handler, err := r.FormFile("avatar")
	if err != nil {
		haveError = true
	}
	if !haveError {
		if handler.Size >= MaxSize {
			SendErrorResponse(w, "Изображение весит больше чем положено, 20mb", 400)
			h.log.Debug("Слишком большой файл")
			return
		}
		defer file.Close()

		if fileType, ok := tools.IsValidType(file); !ok {
			SendErrorResponse(w, fmt.Sprintf("Разрешение %v не поддерживается", fileType), 400)
			h.log.Debug(fmt.Sprintf("Разрешение %v не поддерживается", fileType))
			return
		}

		file.Seek(0, 0)

		filename := tools.GetRandomString()
		if err := storage.UploadFile(file, filename); err != nil {
			SendErrorResponse(w, "Не удалось загрузить файл ", 500)
			h.log.Debug("Не удалось загрузить файл по причине: " + err.Error())
			return
		}

		userConfirm.Avatar = tools.GetStorageUrl(filename)
	}
	desc := r.FormValue("description")
	userConfirm.Description = desc

	if !userConfirm.Handle() {
		SendErrorResponse(w, "Пустые данные, добавлять ничего не будет", http.StatusOK)
		return
	}

	confirmSignUp := userConfirm.Build()

	resp := h.srv.ContinueSignUp(confirmSignUp)
	if resp != nil {
		SendErrorResponse(w, resp.ErrorMessage, resp.ErrorCode)
		h.log.Debug(resp.ErrorMessage)
		return
	}

	SendErrorResponse(w, "", http.StatusOK)
}
