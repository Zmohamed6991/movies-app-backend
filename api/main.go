package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const port = 8080

type application struct {
	DSN string
	Domain string
	DB *sql.DB

}

func main() {
	// set up configuration
	var app application

	// need to load env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading .env file")
		return
	}
	// get the variable
	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = "default_dsn"
	}
	fmt.Println("DSN:", dsn)

	// connect to db
	db, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = db
	defer app.DB.Close()

	app.Domain = "example.com"

	fmt.Println("Server is running on port", port)
	
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}