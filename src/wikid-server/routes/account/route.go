package account

import (
	restful "github.com/emicklei/go-restful"
)

func RegisterRoutes(container *restful.Container) {
	controller := NewController()
	service := &restful.WebService{}
	container.Add(service)

	service.Path("/accounts").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	service.Route(service.POST("/").
		To(controller.PostRoot).
		Reads(&PostRootRequestModel{}).
		Writes(&PostRootResponseModel{}))
}
