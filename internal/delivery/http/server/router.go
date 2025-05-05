package server

import (
	_ "authService/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) createController() {
	s.serv.Use(s.middleware.CORSMiddleware)
	s.serv.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	authGroup := s.serv.Group("/auth")
	authGroup.POST("/login", s.AuthLogin)
	authGroup.POST("/verify-code", s.AuthVerifyCode)
	authGroup.POST("/refresh", s.AuthRefresh)
}
