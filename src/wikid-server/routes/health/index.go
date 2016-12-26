package health

import restful "github.com/emicklei/go-restful"

type GetIndexResponseBody struct {
	Database string `json:"database"`
}

func (this *controller) GetIndex(request *restful.Request, response *restful.Response) {
	const healthy = "healthy"
	const unhealthy = "unhealthy"

	responseBody := &GetIndexResponseBody{
		Database: healthy,
	}

	if err := this.healthService.CheckDatabase(); err != nil {
		responseBody.Database = unhealthy
	}

	response.WriteEntity(responseBody)
}
