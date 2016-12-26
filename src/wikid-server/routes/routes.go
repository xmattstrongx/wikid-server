package routes

import (
	"net/http"
	"wikid-server/routes/accounts"
	"wikid-server/routes/health"

	restful "github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
)

// Build builds the server routes.
func Build() http.Handler {
	container := restful.NewContainer()

	// Register routes.
	accounts.NewController().RegisterRoutes(container)
	health.NewController().RegisterRoutes(container)

	// Build Swagger docs.
	config := swagger.Config{
		WebServices:     container.RegisteredWebServices(),
		ApiPath:         "/swagger.json",
		SwaggerPath:     "/swagger/",
		SwaggerFilePath: "swagger-ui",
	}

	swagger.RegisterSwaggerService(config, container)

	return container
}
