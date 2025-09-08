package routers

import (
	"github.com/Darari17/be-go-tickitz-app/internal/handlers"
	"github.com/Darari17/be-go-tickitz-app/internal/middlewares"
	"github.com/Darari17/be-go-tickitz-app/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func initMovieRouter(router *gin.Engine, db *pgxpool.Pool) {
	movieRepo := repositories.NewMovieRepo(db)
	movieHandler := handlers.NewMovieHandler(movieRepo)

	movieRouter := router.Group("/movies")
	movieRouter.GET("/upcoming", movieHandler.GetUpcomingMovies)
	movieRouter.GET("/popular", movieHandler.GetPopularMovies)
	movieRouter.GET("", movieHandler.GetMoviesWithPagination) // ?limit=10&offset=0
	movieRouter.GET("/:id", movieHandler.GetMovieDetail)
	movieRouter.GET("/:id/schedules", movieHandler.GetSchedule)
	movieRouter.GET("/schedules/:schedule_id/seats", movieHandler.GetAvailableSeats)

	adminMovieRouter := router.Group("/admin/movies", middlewares.VerifyToken, middlewares.AdminOnly)
	adminMovieRouter.GET("", movieHandler.GetAllMovies)
	adminMovieRouter.PUT("/:id", movieHandler.UpdateMovie)
	adminMovieRouter.DELETE("/:id", movieHandler.DeleteMovie)
}
