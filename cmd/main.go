package main

import (
	"log"
	"os"

	"github.com/Darari17/be-go-tickitz-app/internal/config"
	"github.com/Darari17/be-go-tickitz-app/internal/routers"
	"github.com/joho/godotenv"
)

// @title 					Backend Golang Tickitz App
// @version 				1.0
// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization
// @description			RESTful API created using gin for BE GO Tickitz App
// @host						localhost:8080
// @basePath				/
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load env\nCause: ", err.Error())
		return
	}
	log.Println(os.Getenv("DBUSER"))

	db, err := config.InitDB()
	if err != nil {
		log.Println("Failed to connect to database\nCause: ", err.Error())
		return
	}
	defer db.Close()

	if err := config.TestDB(db); err != nil {
		log.Println("Ping to DB failed\nCause: ", err.Error())
		return
	}
	log.Println("DB Connected")

	router := routers.InitRouter(db)
	router.Run("localhost:8080")
}
