package controllers

import (
	"wikid-server/services"

	"wikid-server/models/view"

	restful "github.com/emicklei/go-restful"
)

type IHealthController interface {
	GetRoot(req *restful.Request, res *restful.Response)
}

type healthController struct {
	healthService services.IHealthService
}

// -------------------------------------------------------------------------- //

func NewHealthController() IHealthController {
	return &healthController{
		healthService: services.NewHealthService(),
	}
}

func (this *healthController) GetRoot(req *restful.Request, res *restful.Response) {
	const healthy = "healthy"
	const unhealthy = "unhealthy"

	resBody := &view.HealthGetRootResponse{
		Database: healthy,
	}

	if err := this.healthService.CheckDatabase(); err != nil {
		resBody.Database = unhealthy
	}

	res.WriteEntity(resBody)
}
