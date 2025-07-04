package handler

import (
	_ "simple-file-redirect/docs"
	v1 "simple-file-redirect/internal/app/handler/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterAppRoutes register all routes app
func RegisterAppRoutes(router *gin.Engine) {
	manager := router.Group("/manager")
	// init Swagger.
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1.RegisterV1Routes(manager)
}
