package main

import (
	"fmt"
	"log"
	"net/http"
	"wikid-server/app"
	"wikid-server/repositories"
	"wikid-server/routes"

	restful "github.com/emicklei/go-restful"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	// Get configuration values.
	config, err := app.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to database.
	db, err := sqlx.Open("mysql", config.DbConnStr)
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize repositories.
	repositories.Init(db)

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
