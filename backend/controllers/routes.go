package controllers

import "github.com/ucasers/go-backend/backend/middlewares"

func (s *Server) initializeRoutes() {
	v1 := s.Router.Group("")
	{
		v1.GET("/hello-world", s.HelloWorld)
		v1.POST("/login", s.Login)
		v1.POST("/register", s.Register)
	}

	v2 := s.Router.Group("/user")
	v2.Use(middlewares.TokenAuthMiddleware())
	{
		v2.GET("/getInfo", s.GetUser)
	}

	v3 := s.Router.Group("/extension")
	{
		v3.POST("/upload", middlewares.TokenAuthMiddleware(), s.UploadExtension)
		v3.POST("/modify", middlewares.TokenAuthMiddleware(), s.ModifyExtension)
		v3.POST("/delete", middlewares.TokenAuthMiddleware(), s.DeleteExtension)
		v3.GET("/list", middlewares.TokenAuthMiddleware(), s.ListExtensions)
	}

	v4 := s.Router.Group("/cipherpair")
	{
		v4.POST("/add", middlewares.TokenAuthMiddleware(), s.AddCipherPair)
		v4.POST("/modify", middlewares.TokenAuthMiddleware(), s.ModifyCipherPair)
		v4.POST("/delete", middlewares.TokenAuthMiddleware(), s.DeleteCipherPair)
		v4.GET("/list", middlewares.TokenAuthMiddleware(), s.ListCipherPairs)
		v4.POST("/get-by-title", s.GetExtensionByTitle)
	}
}
