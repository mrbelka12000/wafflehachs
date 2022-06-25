package repository

import (
	"database/sql"
	"go.uber.org/zap"
	"time"
	"wafflehacks/models"
)

type RoomRepo struct {
	db  *sql.DB
	log *zap.SugaredLogger
}

func newRoom(db *sql.DB, zap *zap.SugaredLogger) *RoomRepo {
	return &RoomRepo{db, zap}
}

func (rr *RoomRepo) CreateRoom(room *models.Room) *models.ErrorResponse {

	err := rr.db.QueryRow(`
	INSERT INTO	rooms
	    (clientid, psychoid, expires, isstarted)
	VALUES 
	    ($1, $2,$3,$4)
	RETURNING id
`, room.ClientId, room.PsychoId, time.Now().Add(2*time.Hour), false).Scan(&room.Id)
	if err != nil {
		rr.log.Debug("не удалось создать комнату: " + err.Error())
		return &models.ErrorResponse{ErrorMessage: "Не удалось создать комнату", ErrorCode: 400}
	}
	return nil
}
