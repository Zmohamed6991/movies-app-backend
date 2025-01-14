package dbrepo

import (
	"context"
	"database/sql"
	"time"
	"github.com/zmohamed6991/movies-app-backend/model"
)

type PostgresDBRepo struct {
    DB *sql.DB

}

func (m *PostgresDBRepo) Connection() *sql.DB {
    return m.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) AllMovies() ([]*model.Movies,error) {

    ctx, timeout := context.WithTimeout(context.Background(), dbTimeout) 
    defer timeout()

    query := `
        select
            id, title, release_dte, genre, running_time,
            rating, description, coalesce(image, ''),
        from
            movies
    `
    rows , err := m.DB.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    var movies []*model.Movies

    for rows.Next() {
        var movie model.Movies
        err := rows.Scan(&movie.ID, &movie.Title, &movie.ReleaseDate, &movie.Genre, &movie.RunningTime, &movie.Rating, &movie.Description, &movie.Image)
        if err != nil {
            return nil, err
        }

        movies = append(movies, &movie)
    }  

    return movies, nil
}