package app

import (
	"simple-file-redirect/internal/app/controller"
	"simple-file-redirect/internal/app/service"
	"simple-file-redirect/internal/middleware"
	"simple-file-redirect/internal/storage/converter"
	"simple-file-redirect/internal/storage/manager"

	"github.com/gin-gonic/gin"
)

func RegisterAppRoutes(router *gin.RouterGroup) {
	// Manager Service
	managerSvc := manager.NewService()
	// Storage Service
	converterSvc := converter.NewService()
	// Service
	svc := service.NewService(managerSvc, converterSvc)
	// Controller
	ctrl := controller.NewController(svc)
	group := router.Group("/")
	{
		group.POST("/upload", middleware.AuthMiddleware(), ctrl.UploadArquivo)
		group.GET("/download", middleware.AuthMiddleware(), ctrl.DownloadArquivo)
		group.POST("/convert", middleware.AuthMiddleware(), ctrl.ConvertArquivo)
	}

}
