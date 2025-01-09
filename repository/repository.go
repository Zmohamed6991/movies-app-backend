package repository

import (
	"github.com/zmohamed6991/movies-app-backend/model"
)
type DatabaseRepo interface {
	AllMovies() ([]*model.Movies,error)

}