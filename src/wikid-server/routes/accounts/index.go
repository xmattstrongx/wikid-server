package accounts

import (
	"net/http"
	"time"

	restful "github.com/emicklei/go-restful"
)

type PostIndexRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PostIndexResponseBody struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	CreatedTime time.Time `json:"createdTime"`
}

func (this *controller) PostIndex(request *restful.Request, response *restful.Response) {
	requestBody := &PostIndexRequestBody{}
	if err := request.ReadEntity(requestBody); err != nil {
		response.WriteError(http.StatusBadRequest, err)
	}

	account, err := this.accountService.CreateAccount(requestBody.Email, requestBody.Password)
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
		return
	}

	responseBody := &PostIndexResponseBody{
		ID:          account.ID,
		Email:       account.Email,
		CreatedTime: account.CreatedTime,
	}

	response.WriteHeaderAndEntity(http.StatusCreated, responseBody)
}
