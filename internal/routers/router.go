package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	docs "github.com/Darari17/be-go-tickitz-app/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(db *pgxpool.Pool) *gin.Engine {
	router := gin.Default()

	initAuthRouter(router, db)
	initMovieRouter(router, db)
	initOrderRouter(router, db)
	initProfileRouter(router, db)

	router.Static("/img", "public")

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// router catch all
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Rute Salah",
		})
	})

	return router
}
