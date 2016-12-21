package routes

import (
	"wikid-server/controllers"

	restful "github.com/emicklei/go-restful"
)

func registerAccountRoutes(container *restful.Container) {
	service := &restful.WebService{}
	container.Add(service)

	service.
		Path("/accounts").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	service.Route(service.POST("/").To(controllers.AccountController.PostAccount))
}
