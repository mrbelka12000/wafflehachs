package repository

import (
	"database/sql"
	"fmt"

	"go.uber.org/zap"
)

type ClientRepo struct {
	db  *sql.DB
	log *zap.SugaredLogger
}

func NewClient(db *sql.DB, log *zap.SugaredLogger) *ClientRepo {
	return &ClientRepo{db, log}
}

func (cr *ClientRepo) SignUp() {
	fmt.Println("i was here")
}
