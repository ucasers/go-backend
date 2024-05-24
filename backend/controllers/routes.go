package controllers

func (s *Server) initializeRoutes() {
	v1 := s.Router.Group("/api/v1")
	{
		v1.GET("/users", s.GetUser)
	}
}
