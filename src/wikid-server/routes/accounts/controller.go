package accounts

import (
	"wikid-server/services"

	restful "github.com/emicklei/go-restful"
)

type controller struct {
	accountService services.IAccountService
}

type IController interface {
	RegisterRoutes(container *restful.Container)
	PostIndex(request *restful.Request, response *restful.Response)
}

func NewController() IController {
	return &controller{
		accountService: services.NewAccountService(),
	}
}

// -------------------------------------------------------------------------- //

func (this *controller) RegisterRoutes(container *restful.Container) {
	webService := &restful.WebService{}
	container.Add(webService)

	webService.Path("/accounts").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	webService.Route(webService.POST("/").
		To(this.PostIndex).
		Reads(PostIndexRequestBody{}).
		Writes(PostIndexResponseBody{}))
}
