package handler

import (
	"fmt"
	"net/http"
	"wafflehacks/entities/storage"
	"wafflehacks/tools"
)

func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(500)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		SendErrorResponse(w, "Слишком большой файл", 400)
		h.log.Debug("Файл слишком много весит: " + err.Error())
		return
	}

	file, handler, err := r.FormFile("avatar")
	if err != nil {
		SendErrorResponse(w, "Не удалось получить файл", 400)
		h.log.Debug("Не удалось получить файл по причине: " + err.Error())
		return
	}

	if handler.Size >= 10<<20 {
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
	h.log.Info(tools.GetStorageUrl(filename))
	w.Write([]byte(tools.GetStorageUrl(filename)))
	// http.Redirect(w, r, tools.GetStorageUrl(filename), 301)
}
