package repository

import (
	"database/sql"
	

	"github.com/zmohamed6991/movies-app-backend/model"
)
type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*model.Movies,error)
	
	

}