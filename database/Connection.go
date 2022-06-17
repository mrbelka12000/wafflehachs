package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func GetConnection() (*sql.DB, error) {
	fmt.Print(getConnectionString())
	db, err := sql.Open("postgres", getConnectionString())
	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}

func getConnectionString() string {
	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))
	return connStr
}
