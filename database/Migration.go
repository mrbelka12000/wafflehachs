package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	"go.uber.org/zap"
)

func UpTables(db *sql.DB, log *zap.SugaredLogger) {
	schemaUpDir := os.Getenv("Schema_UP")
	dir, err := ioutil.ReadDir(schemaUpDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range dir {
		body, err := ioutil.ReadFile(schemaUpDir + "/" + file.Name())
		if err != nil {
			log.Fatal("Cant read file: ", err)
		}

		if _, err = db.Exec(string(body)); err != nil {
			log.Info(fmt.Sprintf("Миграция %v не может отработать по причине %v", file.Name(), err.Error()))
		}
	}
}
