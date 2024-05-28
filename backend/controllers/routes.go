package controllers

func (s *Server) initializeRoutes() {
	v1 := s.Router.Group("/api")
	{
		v1.GET("/users", s.GetUser)
		v1.GET("/hello-world", s.HelloWorld)
	}
}
