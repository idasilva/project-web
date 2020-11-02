package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/idasilva/project-web/api/server"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

type API struct {
	context context.Context

	router *mux.Router
	sample  *server.Sample
}

func (a *API) Serve() error {
	ctx, cancel := context.WithCancel(a.context)

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)
		<-ch
		cancel()
		a.sample.Shutdown(ctx)
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer cancel()
		a.sample.Start(a)
	}()
	wg.Wait()

	return nil
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var h http.Handler = a.router
	h.ServeHTTP(w, r)
}

func (a *API) AddRoute(r ...*RouteInfo) {

	for _, route := range r {
		a.router.HandleFunc(route.Path, route.Handler)
		fmt.Println(route.Name, route.Method, route.Path)

	}
}

func NewContext() *API {
	return &API{
		context: context.Background(),
		sample:  server.NewSample(),
		router:  mux.NewRouter().StrictSlash(true),
	}
}
