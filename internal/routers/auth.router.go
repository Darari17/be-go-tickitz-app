package routers

import (
	"github.com/Darari17/be-go-tickitz-app/internal/handlers"
	"github.com/Darari17/be-go-tickitz-app/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func initAuthRouter(router *gin.Engine, db *pgxpool.Pool) {
	authGroup := router.Group("/auth")

	authRepo := repositories.NewAuthRepo(db)
	authHandler := handlers.NewAuthHandler(authRepo)

	authGroup.POST("/login", authHandler.Login)
	authGroup.POST("/register", authHandler.Register)
}
