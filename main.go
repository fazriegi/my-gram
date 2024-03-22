package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fazriegi/my-gram/config"
	"github.com/fazriegi/my-gram/route"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	db := config.NewDatabase()
	logger := config.NewLogger()
	app := config.NewGin()

	routeConfig := route.RouteConfig{
		App:    app,
		DB:     db,
		Logger: logger,
	}
	routeConfig.NewRoute()

	port := os.Getenv("PORT")

	log.Fatal(app.Run(fmt.Sprintf(":%s", port)))
}
