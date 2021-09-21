package database_golang_database

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "maher"
	password = "maiing11"
	dbname   = "youdb"
)

func GetConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
