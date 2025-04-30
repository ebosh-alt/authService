package server

import (
	_ "authService/docs"

	// "github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) createController() {
	s.serv.Use(s.middleware.CORSMiddleware)
	s.serv.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	authGroup := s.serv.Group("/auth")
	authGroup.POST("/login")
	authGroup.POST("/verify-code")
	authGroup.POST("/refresh")
}
