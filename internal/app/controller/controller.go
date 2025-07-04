package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	UploadArquivo(c *gin.Context)
	DownloadArquivo(c *gin.Context)
	ConvertArquivo(c *gin.Context)
	ListenArquivo(c *gin.Context)
}
