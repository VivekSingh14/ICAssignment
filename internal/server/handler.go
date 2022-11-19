package server

import "github.com/shortUrl/ICAssignment/internal/controller"

func (s *Server) registerHandler() {
	controller.RegisterHandlers(s.router, s.config)
}
