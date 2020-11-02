package api

import (
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request)

type RouteInfo struct {
	Handler Handler //func hello(w http.ResponseWriter, r *http.Request)
	Path    string  //http://localhost/helloworld
	Method  string  //get post put delete
	Name    string
}
