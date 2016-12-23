package main

import (
	"fmt"
	"log"
	"net/http"
	"wikid-server/app"
	"wikid-server/routes"

	restful "github.com/emicklei/go-restful"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Get config values.
	config := app.GetConfig()

	// Register routes.
	container := restful.NewContainer()
	routes.Register(container)

	// Configure server.
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: container,
	}

	// Start listening.
	log.Printf("Server configured to listen on port %d.\n", config.Port)
	log.Fatalln(server.ListenAndServe())
}
