package main

import (

	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/zmohamed6991/movies-app-backend/repository"
	"github.com/zmohamed6991/movies-app-backend/repository/dbrepo"
)

const port = 8080

type application struct {
	DSN string
	Domain string
	DB repository.DatabaseRepo
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
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	app.Domain = "example.com"

	fmt.Println("Server is running on port", port)

	fmt.Printf("Type of app: %T\n", app)

	
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}