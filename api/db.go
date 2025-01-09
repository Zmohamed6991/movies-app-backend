package main

import (
	"database/sql"
	"log"

	//"github.com/jackc/pgx/pgconn"
	//"github.com/jackc/pgx/v5"
	//"github.com/jackc/pgx/v5/stdlib"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
		if err != nil {
			return nil, err
		}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	
	return db, nil
}

func (app *application) connectToDB() (*sql.DB, error) {

	connection, err := openDB(app.DSN)
	if err != nil {
		return nil, err
	}

	log.Println("connected to Postgres")

	return connection, nil

}