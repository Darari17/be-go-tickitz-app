package models

import "time"

type Movie struct {
	ID          int        `db:"id" json:"id"`
	Backdrop    string     `db:"backdrop_path" json:"backdrop"`
	Overview    string     `db:"overview" json:"overview"`
	Popularity  float64    `db:"popularity" json:"popularity"`
	Poster      string     `db:"poster_path" json:"poster"`
	ReleaseDate time.Time  `db:"release_date" json:"release_date"`
	Duration    int        `db:"duration" json:"duration"`
	Title       string     `db:"title" json:"title"`
	Director    string     `db:"director_name" json:"director"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
	Genres      []Genre    `db:"-" json:"genres"`
	Casts       []Cast     `db:"-" json:"casts"`
}

type Cast struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type MovieCast struct {
	ID      int `db:"id" json:"id"`
	MovieID int `db:"movies_id" json:"movie_id"`
	CastID  int `db:"casts_id" json:"cast_id"`
}

type Genre struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type MovieGenre struct {
	ID      int `db:"id" json:"id"`
	MovieID int `db:"movies_id" json:"movie_id"`
	GenreID int `db:"genres_id" json:"genre_id"`
}

// untuk admin
type UpdateMovieRequest struct {
	Title       string    `json:"title" binding:"required" example:"Avengers: Endgame"`
	Poster      string    `json:"poster" binding:"required" example:"https://image.tmdb.org/t/p/w500/poster.jpg"`
	Backdrop    string    `json:"backdrop" binding:"required" example:"https://image.tmdb.org/t/p/w500/backdrop.jpg"`
	Overview    string    `json:"overview" binding:"required" example:"After the devastating events of Infinity War..."`
	ReleaseDate time.Time `json:"release_date" binding:"required" example:"2019-04-26T00:00:00Z"`
	Duration    int       `json:"duration" binding:"required" example:"180"`
	Director    string    `json:"director" binding:"required" example:"Anthony Russo, Joe Russo"`
	Popularity  float64   `json:"popularity" binding:"required" example:"95.6"`
}
