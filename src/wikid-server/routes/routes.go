package routes

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
)

// Init initializes application routes.
func Init() http.Handler {
	container := restful.NewContainer()

	registerAPIRoutes(container)

	registerDocRoutes(container)

	return container
}
