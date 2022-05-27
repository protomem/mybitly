package httpserver

import (
	"context"
	"net/http"
	"strconv"
	"time"
)

type HttpServer struct {
	httpServer *http.Server
}

func New(h http.Handler, port int) *HttpServer {
	return &HttpServer{
		httpServer: &http.Server{
			Addr: ":" + strconv.Itoa(port),

			Handler: h,

			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,

			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *HttpServer) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *HttpServer) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
