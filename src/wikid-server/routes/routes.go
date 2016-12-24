package routes

import (
	restful "github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
)

// Register registers the routes with the provided container.
func Register(container *restful.Container) {
	registerAccountRoutes(container)
	registerHealthRoutes(container)

	// Swagger docs must be registered last.
	registerSwaggerDocs(container)
}

func registerSwaggerDocs(container *restful.Container) {
	config := swagger.Config{
		WebServices:     container.RegisteredWebServices(),
		ApiPath:         "/swagger.json",
		SwaggerPath:     "/swagger/",
		SwaggerFilePath: "swagger-ui",
	}

	swagger.RegisterSwaggerService(config, container)
}
