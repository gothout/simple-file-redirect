package server

import (
	"fmt"
	"log"
	"simple-file-redirect/internal/app/handler"
	env "simple-file-redirect/internal/configuration/env/server"

	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	router := gin.Default()
	//register handlers
	handler.RegisterAppRoutes(router)
	return router
}

func StartServer(router *gin.Engine) {
	// start http
	address_http := fmt.Sprintf("%s:%s", env.GetHostServer(), env.GetHTTPPort())
	go func() {
		if err := router.Run(address_http); err != nil {
			log.Fatalf("error inicialize server http: %v", err)
		}
		log.Printf("Server listen on http://%s", address_http)
	}()
	select {}
}
