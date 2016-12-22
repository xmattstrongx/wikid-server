package routes

import (
	"net/http"
	"wikid-server/routes/account"

	restful "github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
)

// Init initializes the server routes.
func Init() http.Handler {
	container := restful.NewContainer()

	account.RegisterRoutes(container)

	// The Swagger routes must be registered last.
	registerSwaggerRoutes(container)

	return container
}

func registerSwaggerRoutes(container *restful.Container) {
	config := swagger.Config{
		WebServices:     container.RegisteredWebServices(),
		ApiPath:         "/swagger.json",
		SwaggerPath:     "/swagger/",
		SwaggerFilePath: "swagger-ui",
	}

	swagger.RegisterSwaggerService(config, container)
}
