package handler

import (
	"fmt"
	"net/http"
	request "wafflehacks/entities/requests"
	"wafflehacks/entities/storage"
	"wafflehacks/models"
	"wafflehacks/tools"
)

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := h.getUserId(r)
	if err != nil {
		SendErrorResponse(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	r.ParseMultipartForm(10 << 20)

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

		if uu.HaveAvater {
			if err = storage.DeleteFile(tools.GetFileNameFromUrl(uu.AvatarUrl)); err != nil {
				SendErrorResponse(w, "Не удалось обновить фотографию", 200)
				h.log.Debug("Не удалось обновить фотографию: " + err.Error())
				return
			}
		}

		filename := tools.GetRandomString()

		if err := storage.UploadFile(file, filename); err != nil {
			SendErrorResponse(w, "Не удалось загрузить файл ", 200)
			h.log.Debug("Не удалось загрузить файл по причине: " + err.Error())
			return
		}

		csu := &models.ContinueSignUp{
			UserID:      user.ID,
			Description: user.Description,
			Avatar:      filename,
		}

		if resp := h.srv.ContinueSignUp(csu); resp != nil {
			SendErrorResponse(w, resp.ErrorMessage, 200)
			h.log.Debug(resp.ErrorMessage)
			return
		}
	}

	SendErrorResponse(w, "", 200)
}