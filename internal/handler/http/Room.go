package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wafflehacks/entities/busymode"
	request "wafflehacks/entities/requests"
	"wafflehacks/entities/response"
	"wafflehacks/models"
)

func (h *Handler) RegisterRoom(w http.ResponseWriter, r *http.Request) {

	cr := &request.CreateRoomRequest{}
	if err := json.NewDecoder(r.Body).Decode(&cr); err != nil {
		response.SendErrorResponse(w, "ошибка дессириализации", 400)
		h.log.Debug("ошибка декодирования: " + err.Error())
		return
	}

	if cr.IsValid() {
		response.SendErrorResponse(w, "Укажите никнейм для создания комнаты", 400)
		h.log.Debug("не указан никнейм")
		return
	}

	user, err := h.getUserId(r)
	if err != nil {
		response.SendErrorResponse(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	psycho, resp := h.srv.GetByUsername(cr.Username)
	if resp != nil {
		response.SendErrorResponse(w, resp.ErrorMessage, resp.ErrorCode)
		h.log.Debug(resp.ErrorMessage)
		return
	}

	if psycho.BusyMode != busymode.ActiveMode {
		response.SendErrorResponse(w, "Психолог в данный момент занят", 400)
		h.log.Debug(fmt.Sprintf("Психолог c ID %v в данный момент занят", psycho.ID))
		return
	}

	room := &models.Room{
		ClientId: user.ID,
		PsychoId: psycho.ID,
	}

	resp = h.srv.CreateRoom(room)
	if resp != nil {
		response.SendErrorResponse(w, resp.ErrorMessage, resp.ErrorCode)
		return
	}

}
