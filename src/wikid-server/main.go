package main

import (
	"fmt"
	"log"
	"net/http"
	"wikid-server/repositories"
	"wikid-server/routes"

	"github.com/caarlos0/env"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Port      int    `env:"PORT"        envDefault:"5000"`
	DBConnStr string `env:"DB_CONN_STR" envDefault:"root:password@tcp(localhost:3306)/wikid"`
}

func main() {
	config := &Config{}
	if err := env.Parse(config); err != nil {
		log.Fatalln(err)
	}

	db, err := sqlx.Open("mysql", config.DBConnStr)
	if err != nil {
		log.Fatalln(err)
	}

	repositories.Init(db)
	routes := routes.Build()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: routes,
	}

	log.Printf("Server configured to listen on port %d.\n", config.Port)
	log.Fatalln(server.ListenAndServe())
}
