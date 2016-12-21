package controllers

import (
	"net/http"

	"wikid-server/services"

	restful "github.com/emicklei/go-restful"
)

type IAuthenticationController interface {
	PostPassword(req *restful.Request, res *restful.Response)
	PostSession(req *restful.Request, res *restful.Response)
}

type authenticationController struct {
	accountService services.IAccountService
}

// -------------------------------------------------------------------------- //

func NewAuthenticationController() IAuthenticationController {
	return &authenticationController{
		accountService: services.NewAccountService(),
	}
}

func (this *authenticationController) PostPassword(req *restful.Request, res *restful.Response) {
	// Parse request body.
	body := &struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := req.ReadEntity(body); err != nil {
		res.WriteError(http.StatusBadRequest, err)
	}

	// Validate credentials.

}

func (this *authenticationController) PostSession(req *restful.Request, res *restful.Response) {

}
