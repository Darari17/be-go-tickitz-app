package repositories

import (
	"context"

	"github.com/Darari17/be-go-tickitz-app/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MovieRepo struct {
	db *pgxpool.Pool
}

func NewMovieRepo(db *pgxpool.Pool) *MovieRepo {
	return &MovieRepo{db: db}
}

// ===========================
// untuk fetch casts dan genres
func (mr *MovieRepo) fetchGenres(ctx context.Context, movieID int) ([]models.Genre, error) {
	rows, err := mr.db.Query(ctx, `
		SELECT g.id, g.name
		FROM genres g
		INNER JOIN movies_genres mg ON g.id = mg.genres_id
		WHERE mg.movies_id = $1
	`, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []models.Genre
	for rows.Next() {
		var g models.Genre
		if err := rows.Scan(&g.ID, &g.Name); err != nil {
			return nil, err
		}
		genres = append(genres, g)
	}
	return genres, nil
}

func (mr *MovieRepo) fetchCasts(ctx context.Context, movieID int) ([]models.Cast, error) {
	rows, err := mr.db.Query(ctx, `
		SELECT c.id, c.name
		FROM casts c
		INNER JOIN movies_casts mc ON c.id = mc.casts_id
		WHERE mc.movies_id = $1
	`, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var casts []models.Cast
	for rows.Next() {
		var c models.Cast
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		casts = append(casts, c)
	}
	return casts, nil
}

//===========================

func (mr *MovieRepo) GetUpcomingMovies(ctx context.Context) ([]models.Movie, error) {
	sql := `
		SELECT id, backdrop_path, overview, popularity, poster_path,
		       release_date, duration, title, director_name,
		       created_at, updated_at
		FROM movies
		WHERE release_date > NOW()
		ORDER BY release_date ASC
	`
	rows, err := mr.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var m models.Movie
		if err := rows.Scan(
			&m.ID, &m.Backdrop, &m.Overview, &m.Popularity, &m.Poster,
			&m.ReleaseDate, &m.Duration, &m.Title, &m.Director,
			&m.CreatedAt, &m.UpdatedAt,
		); err != nil {
			return nil, err
		}

		m.Genres, _ = mr.fetchGenres(ctx, m.ID)
		m.Casts, _ = mr.fetchCasts(ctx, m.ID)

		movies = append(movies, m)
	}
	return movies, nil
}

func (mr *MovieRepo) GetPopularMovies(ctx context.Context) ([]models.Movie, error) {
	sql := `
		SELECT id, backdrop_path, overview, popularity, poster_path,
		       release_date, duration, title, director_name,
		       created_at, updated_at
		FROM movies
		ORDER BY popularity DESC
		LIMIT 10
	`
	rows, err := mr.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var m models.Movie
		if err := rows.Scan(
			&m.ID, &m.Backdrop, &m.Overview, &m.Popularity, &m.Poster,
			&m.ReleaseDate, &m.Duration, &m.Title, &m.Director,
			&m.CreatedAt, &m.UpdatedAt,
		); err != nil {
			return nil, err
		}

		m.Genres, _ = mr.fetchGenres(ctx, m.ID)
		m.Casts, _ = mr.fetchCasts(ctx, m.ID)

		movies = append(movies, m)
	}
	return movies, nil
}

func (mr *MovieRepo) GetMoviesWithPagination(ctx context.Context, limit, offset int) ([]models.Movie, error) {
	sql := `
		SELECT id, backdrop_path, overview, popularity, poster_path,
		       release_date, duration, title, director_name,
		       created_at, updated_at
		FROM movies
		ORDER BY id ASC
		LIMIT $1 OFFSET $2
	`
	rows, err := mr.db.Query(ctx, sql, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var m models.Movie
		if err := rows.Scan(
			&m.ID, &m.Backdrop, &m.Overview, &m.Popularity, &m.Poster,
			&m.ReleaseDate, &m.Duration, &m.Title, &m.Director,
			&m.CreatedAt, &m.UpdatedAt,
		); err != nil {
			return nil, err
		}

		m.Genres, _ = mr.fetchGenres(ctx, m.ID)
		m.Casts, _ = mr.fetchCasts(ctx, m.ID)

		movies = append(movies, m)
	}
	return movies, nil
}

func (mr *MovieRepo) GetSchedule(ctx context.Context, movieID int) ([]models.Schedule, error) {
	sql := `
		SELECT id, movies_id, cinemas_id, times_id, locations_id, date
		FROM schedules
		WHERE movies_id = $1
		ORDER BY date ASC
	`
	rows, err := mr.db.Query(ctx, sql, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []models.Schedule
	for rows.Next() {
		var s models.Schedule
		if err := rows.Scan(&s.ID, &s.MovieID, &s.CinemaID, &s.TimeID, &s.LocationID, &s.Date); err != nil {
			return nil, err
		}
		schedules = append(schedules, s)
	}
	return schedules, nil
}

func (mr *MovieRepo) GetAvailableSeats(ctx context.Context, scheduleID int) ([]models.Seat, error) {
	sql := `
		SELECT s.id, s.seat_code
		FROM seats s
		WHERE s.id NOT IN (
			SELECT os.seats_id
			FROM orders_seats os
			INNER JOIN orders o ON o.id = os.orders_id
			WHERE o.schedules_id = $1
		)
		ORDER BY s.seat_code ASC
	`
	rows, err := mr.db.Query(ctx, sql, scheduleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seats []models.Seat
	for rows.Next() {
		var seat models.Seat
		if err := rows.Scan(&seat.ID, &seat.SeatCode); err != nil {
			return nil, err
		}
		seats = append(seats, seat)
	}
	return seats, nil
}

func (mr *MovieRepo) GetMovieDetail(ctx context.Context, id int) (*models.Movie, error) {
	sql := `
		SELECT id, backdrop_path, overview, popularity, poster_path,
		       release_date, duration, title, director_name,
		       created_at, updated_at
		FROM movies
		WHERE id = $1
	`
	var m models.Movie
	err := mr.db.QueryRow(ctx, sql, id).Scan(
		&m.ID, &m.Backdrop, &m.Overview, &m.Popularity, &m.Poster,
		&m.ReleaseDate, &m.Duration, &m.Title, &m.Director,
		&m.CreatedAt, &m.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	m.Genres, _ = mr.fetchGenres(ctx, m.ID)
	m.Casts, _ = mr.fetchCasts(ctx, m.ID)

	return &m, nil
}

// untuk admin ==============

func (mr *MovieRepo) GetAllMovies(ctx context.Context) ([]models.Movie, error) {
	sql := `
		SELECT id, backdrop_path, overview, popularity, poster_path,
		       release_date, duration, title, director_name,
		       created_at, updated_at
		FROM movies
		ORDER BY created_at DESC
	`
	rows, err := mr.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var m models.Movie
		if err := rows.Scan(
			&m.ID, &m.Backdrop, &m.Overview, &m.Popularity, &m.Poster,
			&m.ReleaseDate, &m.Duration, &m.Title, &m.Director,
			&m.CreatedAt, &m.UpdatedAt,
		); err != nil {
			return nil, err
		}

		m.Genres, _ = mr.fetchGenres(ctx, m.ID)
		m.Casts, _ = mr.fetchCasts(ctx, m.ID)

		movies = append(movies, m)
	}
	return movies, nil
}

func (mr *MovieRepo) DeleteMovie(ctx context.Context, id int) error {
	_, err := mr.db.Exec(ctx, `DELETE FROM movies_genres WHERE movies_id=$1`, id)
	if err != nil {
		return err
	}
	_, err = mr.db.Exec(ctx, `DELETE FROM movies_casts WHERE movies_id=$1`, id)
	if err != nil {
		return err
	}
	_, err = mr.db.Exec(ctx, `DELETE FROM schedules WHERE movies_id=$1`, id)
	if err != nil {
		return err
	}
	_, err = mr.db.Exec(ctx, `DELETE FROM movies WHERE id=$1`, id)
	return err
}

func (mr *MovieRepo) UpdateMovie(ctx context.Context, movie models.Movie) error {
	sql := `
		UPDATE movies
		SET title=$1, poster_path=$2, backdrop_path=$3, overview=$4,
		    release_date=$5, duration=$6, director_name=$7, popularity=$8,
		    updated_at=NOW()
		WHERE id=$9
	`
	_, err := mr.db.Exec(ctx, sql,
		movie.Title, movie.Poster, movie.Backdrop, movie.Overview,
		movie.ReleaseDate, movie.Duration, movie.Director, movie.Popularity,
		movie.ID,
	)
	return err
}
