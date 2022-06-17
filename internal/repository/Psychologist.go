package repository

import (
	"database/sql"
	"fmt"

	"go.uber.org/zap"
)

type PsychologystRepo struct {
	db  *sql.DB
	log *zap.SugaredLogger
}

func NewPsycho(db *sql.DB, log *zap.SugaredLogger) *PsychologystRepo {
	return &PsychologystRepo{db, log}
}

func (pr *PsychologystRepo) SignUp() {
	fmt.Println("i was here")
}
