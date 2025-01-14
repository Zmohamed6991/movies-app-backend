package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	

}


func (app *application) allMovies(w http.ResponseWriter, r *http.Request) {

	movies, err := app.DB.AllMovies()
	if err != nil {
		fmt.Println(nil)
	}

	out, err := json.Marshal(movies)
	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

	
}
