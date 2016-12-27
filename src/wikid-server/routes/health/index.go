package health

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
)

type GetIndexResponseBody struct {
	DatabaseHealthy bool `json:"databaseHealthy"`
}

func (this *controller) GetIndex(request *restful.Request, response *restful.Response) {
	report, err := this.healthService.GenerateReport()
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	responseBody := &GetIndexResponseBody{
		DatabaseHealthy: report.DatabaseHealthy,
	}

	response.WriteEntity(responseBody)
}
