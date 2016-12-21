package controllers

import (
	"net/http"
	"wikid-server/services"

	restful "github.com/emicklei/go-restful"
)

type IAccountController interface {
	PostAccount(req *restful.Request, res *restful.Response)
}

type accountController struct {
	accountService services.IAccountService
	sessionService services.ISessionService
}

// -------------------------------------------------------------------------- //

func NewAccountController() IAccountController {
	return &accountController{
		accountService: services.AccountService,
		sessionService: services.SessionService,
	}
}

func (this *accountController) PostAccount(req *restful.Request, res *restful.Response) {
	// Parse request body.
	body := &struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := req.ReadEntity(body); err != nil {
		res.WriteError(http.StatusBadRequest, err)
	}

	// Create account.
	account, err := this.accountService.CreateAccount(body.Email, body.Password)
	if err != nil {
		// TODO: Return a better error message.
		res.WriteError(http.StatusBadRequest, err)
		return
	}

	// TODO: Create session.
	// session, err := this.sessionService.CreateSession(account.Id)
	// if err != nil {
	// 	// TODO: Return a better error message.
	// 	res.WriteError(http.StatusBadRequest, err)
	// 	return
	// }

	// TODO: Return session.
	res.WriteAsJson(account)
}
