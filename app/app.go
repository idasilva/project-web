package app

import (
	"github.com/idasilva/project-web/api"
	"github.com/idasilva/project-web/app/page"
	"net/http"
)

func Run() *api.API {

	app := api.NewContext()
	app.AddRoute(&api.RouteInfo{
		Handler: page.Index,
		Path:    "/index",
		Method:  http.MethodGet,
		Name:    "Initial page",
	},&api.RouteInfo{
		Handler: page.Form,
		Path:    "/form/1",
		Method:  http.MethodGet,
		Name:    "Form",
	},&api.RouteInfo{
		Handler: page.Register,
		Path:    "/register/2",
		Method:  http.MethodPost,
		Name:    "Register",
	})

	return app
}
