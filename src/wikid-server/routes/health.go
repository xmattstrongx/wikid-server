package routes

import (
	"wikid-server/controllers"
	"wikid-server/models/view"

	restful "github.com/emicklei/go-restful"
)

func registerHealthRoutes(container *restful.Container) {
	healthController := controllers.NewHealthController()
	service := &restful.WebService{}
	container.Add(service)

	service.Path("/health").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	service.Route(service.GET("/").
		To(healthController.GetRoot).
		Writes(view.HealthGetRootResponse{}))
}
