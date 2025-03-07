package server

import (
	_ "authSerivce/docs"

	// "github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) createController() {
	s.serv.Use(s.middleware.CORSMiddleware)
	s.serv.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//currencyGroup := s.serv.Group("/currency")
	//currencyGroup.POST("/add", s.CurrencyAdd)
	//currencyGroup.POST("/remove", s.CurrencyRemove)
	//currencyGroup.POST("/price", s.CurrencyPrice)

}
