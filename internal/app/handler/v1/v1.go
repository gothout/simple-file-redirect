package v1

import (
	"simple-file-redirect/internal/app/handler/v1/app"

	"github.com/gin-gonic/gin"
)

func RegisterV1Routes(router *gin.RouterGroup) {
	v1Group := router.Group("/v1")
	app.RegisterAppRoutes(v1Group)
}
