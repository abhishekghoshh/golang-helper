package server

import (
	"di/logger"
	"net/http"
)

type Server struct {
	Logger logger.Logger
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Logger.Log("received request from " + r.Host)
	s.Logger.Log("requested path is " + r.URL.Path)
}
