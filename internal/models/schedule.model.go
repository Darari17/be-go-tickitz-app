package models

import "time"

type Schedule struct {
	ID         int       `db:"id" json:"id"`
	MovieID    int       `db:"movies_id" json:"movie_id"`
	CinemaID   int       `db:"cinemas_id" json:"cinema_id"`
	TimeID     int       `db:"times_id" json:"time_id"`
	LocationID int       `db:"locations_id" json:"location_id"`
	Date       time.Time `db:"date" json:"date"`
}

type Cinema struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Location struct {
	ID       int    `db:"id" json:"id"`
	Location string `db:"location" json:"location"`
}

type Time struct {
	ID   int    `db:"id" json:"id"`
	Time string `db:"time" json:"time"`
}
