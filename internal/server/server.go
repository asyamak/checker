package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"restapi/config"
)

type Server struct {
	Srv    *http.Server
	notify chan error
}

func New(cnf *config.Config, router *http.ServeMux) *Server {
	return &Server{
		Srv: &http.Server{
			Addr:           ":" + cnf.Port,
			Handler:        router,
			MaxHeaderBytes: cnf.MaxHeaderBytes,
			ReadTimeout:    time.Duration(cnf.ReadTimeOut * int(time.Second)),
			WriteTimeout:   time.Duration(cnf.WriteTimeOut * int(time.Second)),
		},
		notify: make(chan error),
	}
}

func (s *Server) Start() {
	go func() {
		log.Printf("server initialized http://localhost%v\n", s.Srv.Addr)
		s.notify <- s.Srv.ListenAndServe()
		// close(s.notify)
		// log.Printf("closing channel: %v\n", s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Srv.Shutdown(ctx)
}
