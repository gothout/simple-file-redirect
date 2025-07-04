package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	UploadArquivo(c *gin.Context)
}
