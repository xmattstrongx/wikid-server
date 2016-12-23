package controllers

import (
	"net/http"
	"wikid-server/models/view"
	"wikid-server/services"

	restful "github.com/emicklei/go-restful"
)

type IAccountController interface {
	PostRoot(req *restful.Request, res *restful.Response)
}

type accountController struct {
	accountService services.IAccountService
}

var NewAccountController = func() IAccountController {
	return &accountController{
		accountService: services.NewAccountService(),
	}
}

// -------------------------------------------------------------------------- //

func (this *accountController) PostRoot(req *restful.Request, res *restful.Response) {
	// Parse request body.
	reqBody := &view.AccountPostRootRequest{}
	if err := req.ReadEntity(reqBody); err != nil {
		res.WriteError(http.StatusBadRequest, err)
	}

	// Create account.
	account, err := this.accountService.CreateAccount(reqBody.Email, reqBody.Password)
	if err != nil {
		res.WriteError(http.StatusBadRequest, err)
		return
	}

	// Return account.
	resBody := &view.AccountPostRootResponse{
		Id:          account.Id,
		Email:       account.Email,
		CreatedTime: account.CreatedTime,
	}

	res.WriteHeaderAndEntity(http.StatusCreated, resBody)
}
