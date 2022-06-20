package repository

import (
	"database/sql"
	"time"
	"wafflehacks/tools"

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

func (r *SessionRepo) GetUserByCookie(cookie string) (*models.User, *models.ErrorResponse) {
	user := &models.User{}
	avatar := sql.NullString{}
	desc := sql.NullString{}
	err := r.db.QueryRow(`
	SELECT 
	    users.id, users.firstname, users.lastname, users.username, users.avatarurl, users.age, users.description
	FROM
	    session
	JOIN 
	    users on users.id = session.userid
	WHERE 
	    uuid=$1
`, cookie).Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Username, &avatar, &user.Age, &desc)
	if err != nil {
		r.log.Debug("Не удалось найти пользователя: " + err.Error())
		return nil, &models.ErrorResponse{ErrorMessage: "Пользователь не найден", ErrorCode: 400}
	}

	if user.ID == 0 {
		r.log.Debug("Не удалось найти пользователя: " + err.Error())
		return nil, &models.ErrorResponse{ErrorMessage: "Пользователь не найден", ErrorCode: 400}
	}

	user.AvatarUrl = tools.GetStorageUrl(avatar.String)
	user.Description = desc.String
	return user, nil
}
