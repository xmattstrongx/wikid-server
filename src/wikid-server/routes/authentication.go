package routes

import (
	"wikid-server/controllers"

	restful "github.com/emicklei/go-restful"
)

func registerAuthenticationRoutes(container *restful.Container) {
	authenticationController := controllers.NewAuthenticationController()
	service := &restful.WebService{}
	container.Add(service)

	service.
		Path("/authentication").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	service.Route(service.POST("/password").To(authenticationController.PostPassword))
	service.Route(service.POST("/session").To(authenticationController.PostSession))
}
