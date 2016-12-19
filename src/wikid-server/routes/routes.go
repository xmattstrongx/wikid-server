package routes

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
)

// Init initializes application routes.
func Init() http.Handler {
	ws := &restful.WebService{}
	ws.Route(ws.GET("/").To(func(req *restful.Request, res *restful.Response) {
		res.Write([]byte("hello world"))
	}))

	container := restful.NewContainer()
	container.Add(ws)

	return container
}
