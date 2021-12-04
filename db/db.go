package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	HOST = "database"
	PORT = 5432
)

type Database struct {
	Connection *sql.DB
}

func Initialize(username, password, database string) (Database, error) {
	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}

	db.Connection = conn
	err = db.Connection.Ping()

	if err != nil {
		return db, err
	}

	return db, nil
}