package repository

import (
	"database/sql"
	"time"

	"wafflehacks/models"

	"go.uber.org/zap"
)

type SessionRepo struct {
	db  *sql.DB
	log *zap.SugaredLogger
}

func newSession(db *sql.DB, log *zap.SugaredLogger) *SessionRepo {
	return &SessionRepo{
		db:  db,
		log: log,
	}
}

func (r *SessionRepo) CreateSession(session *models.SessionResponse) *models.ErrorResponse {
	_, err := r.db.Exec(`
		INSERT INTO session
			(userid, uuid, expires_at)
		VALUES
			($1, $2, $3)
	`, session.ID, session.Cookie, time.Now().Add(10*time.Minute))
	if err != nil {
		r.log.Debug("inserting cookie: " + err.Error())
		return &models.ErrorResponse{ErrorMessage: "insert cookie", ErrorCode: 500}
	}

	return nil
}

func (r *SessionRepo) GetUserIdByCookie(cookie string) (int, *models.ErrorResponse) {
	id := 0
	err := r.db.QueryRow(`
	SELECT 
	    userid
	FROM
	    session
	WHERE 
	    cookie=$1
`, cookie).Scan(&id)
	if err != nil {
		r.log.Debug("Не удалось найти пользователя: " + err.Error())
		return 0, &models.ErrorResponse{ErrorMessage: "Пользователь не найден", ErrorCode: 400}
	}
	if id == 0 {
		r.log.Debug("Не удалось найти пользователя: " + err.Error())
		return 0, &models.ErrorResponse{ErrorMessage: "Пользователь не найден", ErrorCode: 400}
	}
	return id, nil
}
