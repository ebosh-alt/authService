package server

import (
	// swaggerFiles "github.com/swaggo/files" // swagger embed files
	_ "authSerivce/docs"

	// "github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Crypto Parser
// @version 1.0
// @description API Server for collecting, storing and displaying the value of cryptocurrencies

// @host localhost:3000
// @BasePath /

func (s *Server) createController() {
	s.serv.Use(s.middleware.CORSMiddleware)
	s.serv.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//currencyGroup := s.serv.Group("/currency")
	//currencyGroup.POST("/add", s.CurrencyAdd)
	//currencyGroup.POST("/remove", s.CurrencyRemove)
	//currencyGroup.POST("/price", s.CurrencyPrice)

}
