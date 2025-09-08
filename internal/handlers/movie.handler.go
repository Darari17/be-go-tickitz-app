package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Darari17/be-go-tickitz-app/internal/models"
	"github.com/Darari17/be-go-tickitz-app/internal/repositories"
	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	movieRepo *repositories.MovieRepo
}

func NewMovieHandler(movieRepo *repositories.MovieRepo) *MovieHandler {
	return &MovieHandler{movieRepo: movieRepo}
}

// GetUpcomingMovies godoc
// @Summary     Get Upcoming Movies
// @Description Upcoming Movies
// @Tags        Movies
// @Produce     json
// @Router      /movies/upcoming [get]
func (mh *MovieHandler) GetUpcomingMovies(ctx *gin.Context) {
	movies, err := mh.movieRepo.GetUpcomingMovies(ctx)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch upcoming movies"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": movies})
}

// GetPopularMovies godoc
// @Summary     Get Popular Movies
// @Description Popular Movies
// @Tags        Movies
// @Produce     json
// @Router      /movies/popular [get]
func (mh *MovieHandler) GetPopularMovies(ctx *gin.Context) {
	movies, err := mh.movieRepo.GetPopularMovies(ctx)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch popular movies"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": movies})
}

// GetMoviesWithPagination godoc
// @Summary     Get Movies with Pagination and Search
// @Description Ambil daftar film dengan pagination dan pencarian berdasarkan judul
// @Tags        Movies
// @Produce     json
// @Param       page      query int    false "Halaman (Default: 1)"
// @Param       pagesize  query int    false "Jumlah data per halaman (Default: 10)"
// @Param       search    query string false "Cari berdasarkan judul film"
// @Router      /movies [get]
func (mh *MovieHandler) GetMoviesWithPagination(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	search := ctx.DefaultQuery("search", "")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	movies, err := mh.movieRepo.GetMoviesWithPagination(ctx, page, pageSize, search)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch movies"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"page":     page,
		"pageSize": pageSize,
		"search":   search,
		"data":     movies,
	})
}

// GetSchedule godoc
// @Summary     Get Schedule by Movie ID
// @Description Get Schedule by Movie ID
// @Tags        Movies
// @Produce     json
// @Param       id path int true "Movie ID"
// @Router      /movies/{id}/schedules [get]
func (mh *MovieHandler) GetSchedule(ctx *gin.Context) {
	movieIDStr := ctx.Param("id")
	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid movie id"})
		return
	}

	schedules, err := mh.movieRepo.GetSchedule(ctx, movieID)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch schedule"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": schedules})
}

// GetAvailableSeats godoc
// @Summary     Get Available Seats
// @Description kursi kosong berdasarkan schedule ID
// @Tags        Movies
// @Produce     json
// @Param       schedule_id path int true "Schedule ID"
// @Router      /movies/schedules/{schedule_id}/seats [get]
func (mh *MovieHandler) GetAvailableSeats(ctx *gin.Context) {
	scheduleIDStr := ctx.Param("schedule_id")
	scheduleID, err := strconv.Atoi(scheduleIDStr)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid schedule id"})
		return
	}

	seats, err := mh.movieRepo.GetAvailableSeats(ctx, scheduleID)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch available seats"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": seats})
}

// GetMovieDetail godoc
// @Summary     Get Movie Detail
// @Description Detail lengkap movie berdasarkan ID
// @Tags        Movies
// @Produce     json
// @Param       id path int true "Movie ID"
// @Router      /movies/{id} [get]
func (mh *MovieHandler) GetMovieDetail(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid movie id"})
		return
	}

	movie, err := mh.movieRepo.GetMovieDetail(ctx, id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{"message": "movie not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": movie})
}

// GetAllMovies godoc
// @Summary     Get All Movies (Admin)
// @Description Semua data Movie untuk admin
// @Tags        Admin-Movies
// @Security    BearerToken
// @Produce     json
// @Router      /admin/movies [get]
func (mh *MovieHandler) GetAllMovies(ctx *gin.Context) {
	movies, err := mh.movieRepo.GetAllMovies(ctx)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch movies"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": movies})
}

// DeleteMovie godoc
// @Summary     Delete Movie (Admin)
// @Description Hapus movie berdasarkan ID
// @Tags        Admin-Movies
// @Security    BearerToken
// @Param       id path int true "Movie ID"
// @Router      /admin/movies/{id} [delete]
func (mh *MovieHandler) DeleteMovie(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid movie id"})
		return
	}

	if err := mh.movieRepo.DeleteMovie(ctx, id); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete movie"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "movie deleted"})
}

// UpdateMovie godoc
// @Summary     Update Movie (Admin)
// @Description Update data movie berdasarkan ID
// @Tags        Admin-Movies
// @Security    BearerToken
// @Accept      json
// @Produce     json
// @Param       id path int true "Movie ID"
// @Param       movie body models.UpdateMovieRequest true "Movie Update Data"
// @Router      /admin/movies/{id} [put]
func (mh *MovieHandler) UpdateMovie(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid movie id"})
		return
	}

	var req models.UpdateMovieRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
		return
	}

	movie := models.Movie{
		ID:          id,
		Title:       req.Title,
		Poster:      req.Poster,
		Backdrop:    req.Backdrop,
		Overview:    req.Overview,
		ReleaseDate: req.ReleaseDate,
		Duration:    req.Duration,
		Director:    req.Director,
		Popularity:  req.Popularity,
	}

	if err := mh.movieRepo.UpdateMovie(ctx, movie); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update movie"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "movie updated"})
}
