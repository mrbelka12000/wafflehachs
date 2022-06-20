package database

import (
	"database/sql"
	"go.uber.org/zap"
	"time"
)

func DeleteExpiredCookie(conn *sql.DB, log *zap.SugaredLogger, ch chan bool) {
	for {
		_, err := conn.Exec(`
		DELETE FROM 
		        session
		WHERE 
		      	expires_at <now()`)
		if err != nil {
			log.Debug(err.Error())
			ch <- true
		}
		time.Sleep(5 * time.Second)
	}
}
