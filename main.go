// @title Simple File Redirect API
// @version 1.0
// @description API para upload, download e conversão de arquivos mp3 para ogg
// @BasePath /

// 🔐 Definição do Bearer Token
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Forneça o token no formato: Bearer <token>

package main

import (
	"log"
	"simple-file-redirect/cmd/server"
	"simple-file-redirect/internal/configuration/env"
)

func main() {
	// Check envs
	err := env.CheckEnvs()
	if err != nil {
		log.Fatal(err)
	}
	// Start server
	router := server.InitServer()
	server.StartServer(router)
}
