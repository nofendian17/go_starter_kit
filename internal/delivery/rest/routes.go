package rest

func (s *Server) routes() {
	// health routes
	s.router.HandleFunc("GET /health", s.handler.Health.Health())
	s.router.HandleFunc("GET /readiness", s.handler.Health.Readiness())
	s.router.HandleFunc("GET /ping", s.handler.Health.Ping())
}
