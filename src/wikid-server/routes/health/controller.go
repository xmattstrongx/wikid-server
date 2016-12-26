package health

import (
	"wikid-server/services"

	restful "github.com/emicklei/go-restful"
)

type controller struct {
	healthService services.IHealthService
}

type IController interface {
	RegisterRoutes(container *restful.Container)
	GetIndex(request *restful.Request, response *restful.Response)
}

func NewController() IController {
	return &controller{
		healthService: services.NewHealthService(),
	}
}

// -------------------------------------------------------------------------- //

func (this *controller) RegisterRoutes(container *restful.Container) {
	webService := &restful.WebService{}
	container.Add(webService)

	webService.Path("/health").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	webService.Route(webService.GET("/").
		To(this.GetIndex).
		Writes(GetIndexResponseBody{}))
}
