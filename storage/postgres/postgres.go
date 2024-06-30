package postgres

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "Garden_Management_service"
	password = "03212164"
)

func ConnectDB() (*sql.DB, error) {
	conn := fmt.Sprintf("localhost=%s port=%d user=%s dbname=%s password=%s ?sslmode=disable",
		host, port, user, dbname, password)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
