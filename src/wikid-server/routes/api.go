package routes

import restful "github.com/emicklei/go-restful"

func registerAPIRoutes(container *restful.Container) {
	service := &restful.WebService{}
	container.Add(service)

	service.
		Path("/api").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	// Account routes.
	service.Route(service.GET("/accounts/{account_id}").To(getAccounts))
	service.Route(service.POST("/accounts").To(postAccounts))
	service.Route(service.PUT("/accounts/{account_id}").To(putAccounts))
	service.Route(service.DELETE("/accounts/{account_id}").To(deleteAccounts))

	// Session routes.
	service.Route(service.GET("/sessions/{session_id}").To(getSessions))
	service.Route(service.POST("/sessions").To(postSessions))
	service.Route(service.DELETE("/sessions/{session_id}").To(deleteSessions))
}

// -- Accounts -------------------------------------------------------------- //

func getAccounts(req *restful.Request, res *restful.Response) {

}

func postAccounts(req *restful.Request, res *restful.Response) {

}

func putAccounts(req *restful.Request, res *restful.Response) {

}

func deleteAccounts(req *restful.Request, res *restful.Response) {

}

// -- Sessions -------------------------------------------------------------- //

func getSessions(req *restful.Request, res *restful.Response) {

}

func postSessions(req *restful.Request, res *restful.Response) {

}

func deleteSessions(req *restful.Request, res *restful.Response) {

}
