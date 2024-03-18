package main

import (
	"fmt"
	"log"

	"github.com/fazriegi/my-gram/config"
	"github.com/fazriegi/my-gram/route"
)

func main() {
	viperConfig := config.NewViper()
	db := config.NewDatabase(viperConfig)
	logger := config.NewLogger(viperConfig)
	app := config.NewGin()

	routeConfig := route.RouteConfig{
		App:    app,
		DB:     db,
		Logger: logger,
	}
	routeConfig.NewRoute()

	port := viperConfig.GetInt("web.port")

	log.Fatal(app.Run(fmt.Sprintf(":%d", port)))
}
