package model

import (
    "time"
)

type Movies struct {
    ID int `json:"id"`
    Title string `json:"title"`
    ReleaseDate int `json:"release_date"`
    Genre string `json:"genre"`
    RunningTime int `json:"running_time"`
    Rating int `json:"rating"`
    Description string `json:"description"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`


}