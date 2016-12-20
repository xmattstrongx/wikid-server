package main

import (
	"fmt"
	"log"
	"net/http"
	"wikid-server/app"
	"wikid-server/repositories"
	"wikid-server/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	// Connect to database.
	db, err := sqlx.Connect("mysql", app.Config.DBConnStr)
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize repositories and routes.
	repositories.Init(db)
	handler := routes.Init()

	// Configure server.
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Config.Port),
		Handler: handler,
	}

	// Start listening.
	log.Printf("Server configured to listen on port %d.\n", app.Config.Port)
	log.Fatalln(server.ListenAndServe())
}
