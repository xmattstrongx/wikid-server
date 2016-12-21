package routes

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
)

// Init initializes the server routes.
func Init() http.Handler {
	container := restful.NewContainer()

	registerAccountRoutes(container)
	registerAuthenticationRoutes(container)

	// The Swagger routes must be registered last.
	registerSwaggerRoutes(container)

	return container
}
