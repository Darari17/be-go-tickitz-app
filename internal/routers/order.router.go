package routers

import (
	"github.com/Darari17/be-go-tickitz-app/internal/handlers"
	"github.com/Darari17/be-go-tickitz-app/internal/middlewares"
	"github.com/Darari17/be-go-tickitz-app/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func initOrderRouter(router *gin.Engine, db *pgxpool.Pool) {
	orderGroup := router.Group("/orders", middlewares.VerifyToken)

	orderRepo := repositories.NewOrderRepo(db)
	orderHandler := handlers.NewOrderHandler(orderRepo)

	orderGroup.POST("", orderHandler.CreateOrder)
	orderGroup.GET("/:id", orderHandler.GetOrderByID)
	orderGroup.GET("/user/:user_id", orderHandler.GetOrdersByUser)
}
