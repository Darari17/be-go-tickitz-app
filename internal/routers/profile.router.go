package routers

import (
	"github.com/Darari17/be-go-tickitz-app/internal/handlers"
	"github.com/Darari17/be-go-tickitz-app/internal/middlewares"
	"github.com/Darari17/be-go-tickitz-app/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func initProfileRouter(router *gin.Engine, db *pgxpool.Pool) {
	profileGroup := router.Group("/profile", middlewares.VerifyToken)

	profileRepo := repositories.NewProfileRepo(db)
	profileHandler := handlers.NewProfileHandler(profileRepo)

	profileGroup.GET("", profileHandler.GetProfile)
	profileGroup.PUT("", profileHandler.UpdateProfile)
}
