package postgres

import (
	"database/sql"
	"fmt"
	
)

func NewDB(host, port, user, password, dbName string) (*sql.DB, error) {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")
	return db, nil
}

// TODO: add database operationss here