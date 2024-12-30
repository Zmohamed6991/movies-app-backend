package main

import (
	"backend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status: "Active",
		Message: "Movflix is running",
		Version: "1.0",
	}

	response, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}


func (app *application) allMovies(w http.ResponseWriter, r *http.Request) {

	fmt.Println(w, "List of movies")
	
	var movies []models.Movies

	ThinkLikeAMan := models.Movies{
		ID: 1,
		Title: "Think like a man",
		ReleaseDate: 2012,
		Genre: "Comedy",
		RunningTime: 122,
		Rating: 7,
		Description: "Four friends conspire to turn the tables on their women when they discover the ladies have been using Steve Harvey's relationship advice against them.",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

		movies = append(movies, ThinkLikeAMan)

		response, err := json.Marshal(movies)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)

	bmh := models.Movies{
		ID: 2,
		Title: "Big Momma House 2",
		ReleaseDate: 2006,
		Genre: "Comedy",
		RunningTime: 99,
		Rating: 5,
		Description: "An FBI agent reprises his disguise as a corpulent old lady and takes a job as a nanny in a crime suspect's house.",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	movies = append(movies, bmh)

	response, err = json.Marshal(movies)
	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)



	
}
