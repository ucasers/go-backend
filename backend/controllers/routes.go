package controllers

import "github.com/ucasers/go-backend/backend/middlewares"

func (s *Server) initializeRoutes() {
	v1 := s.Router.Group("")
	{
		v1.GET("/users", s.GetUser)
		v1.GET("/hello-world", s.HelloWorld)
		v1.POST("/login", s.Login)
		v1.POST("/register", s.Register)
		v1.GET("/get-user", middlewares.TokenAuthMiddleware(), s.GetUser)
	}
}
