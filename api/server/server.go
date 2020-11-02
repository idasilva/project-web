package server

import (
	"fmt"
	"net/http"
	"time"
)

const port = ":8000"

type Sample struct {
	*http.Server
}

func (s *Sample) addr() string {
	s.Addr = port
	return s.Addr

}
func (s *Sample) Start(handler http.Handler) {
	s.Handler = handler
	fmt.Println("Create a new sample server, start at port:", s.addr())
	s.ListenAndServe()

}

func NewSample() *Sample {
	return &Sample{
		Server: &http.Server{
			WriteTimeout: 600 * time.Second,
			ReadTimeout:  600 * time.Second,
		},
	}
}
