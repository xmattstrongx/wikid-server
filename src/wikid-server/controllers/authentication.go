package controllers

import restful "github.com/emicklei/go-restful"

type IAuthenticationController interface {
	PostPassword(req *restful.Request, res *restful.Response)
	PostSession(req *restful.Request, res *restful.Response)
}

type authenticationController struct{}

// -------------------------------------------------------------------------- //

func NewAuthenticationController() IAuthenticationController {
	return &authenticationController{}
}

func (this *authenticationController) PostPassword(req *restful.Request, res *restful.Response) {

}

func (this *authenticationController) PostSession(req *restful.Request, res *restful.Response) {

}
