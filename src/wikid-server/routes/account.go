package routes

import (
	"wikid-server/controllers"
	"wikid-server/models/view"

	restful "github.com/emicklei/go-restful"
)

func registerAccountRoutes(container *restful.Container) {
	accountController := controllers.NewAccountController()
	service := &restful.WebService{}
	container.Add(service)

	service.Path("/accounts").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	service.Route(service.POST("/").
		To(accountController.PostRoot).
		Reads(view.AccountPostRootRequest{}).
		Writes(view.AccountPostRootResponse{}))
}
