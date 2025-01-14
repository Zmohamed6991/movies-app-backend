package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	_"github.com/jackc/pgx/v5"
	_"github.com/jackc/pgx/v5/stdlib"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
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

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Get("/", app.Home)
	mux.Get("/movies", app.allMovies)

	return mux
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