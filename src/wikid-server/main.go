package main

import (
	"fmt"
	"log"
	"net/http"
	"wikid-server/app"
	"wikid-server/routes"
)

func main() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Config.Port),
		Handler: routes.Init(),
	}

	log.Printf("Server configured to listen on port %d.\n", app.Config.Port)
	log.Fatalln(server.ListenAndServe())
}
