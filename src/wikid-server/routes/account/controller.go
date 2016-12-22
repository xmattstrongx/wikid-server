package account

import (
	"net/http"
	"wikid-server/services"

	restful "github.com/emicklei/go-restful"
)

type IController interface {
	PostRoot(req *restful.Request, res *restful.Response)
}

type controller struct {
	accountService services.IAccountService
}

var NewController = func() IController {
	return &controller{
		accountService: services.NewAccountService(),
	}
}

// -------------------------------------------------------------------------- //

func (this *controller) PostRoot(req *restful.Request, res *restful.Response) {
	// Parse request body.
	reqBody := &PostRootRequestModel{}
	if err := req.ReadEntity(reqBody); err != nil {
		res.WriteError(http.StatusBadRequest, err)
	}

	// Create account.
	account, err := this.accountService.CreateAccount(reqBody.Email, reqBody.Password)
	if err != nil {
		res.WriteError(http.StatusBadRequest, err)
		return
	}

	// // Create session.
	// session, err := this.accountService.SignIn(body.Email, body.Password)
	// if err != nil {
	// 	res.WriteError(http.StatusBadRequest, err)
	// 	return
	// }

	// Return session.
	resBody := &PostRootResponseModel{
		SessionId: "session id",
		AccountId: account.Id,
	}

	res.WriteHeaderAndEntity(http.StatusCreated, resBody)
}
