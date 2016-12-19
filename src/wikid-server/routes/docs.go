package routes

import (
	restful "github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
)

func registerDocRoutes(container *restful.Container) {
	config := swagger.Config{
		WebServices:     container.RegisteredWebServices(),
		ApiPath:         "/swagger.json",
		SwaggerPath:     "/docs",
		SwaggerFilePath: "swagger-ui",
	}

	swagger.RegisterSwaggerService(config, container)
}
