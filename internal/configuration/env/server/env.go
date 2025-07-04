package server

import (
	"fmt"
	"os"
)

// Host para o servidor
func GetHostServer() string {
	return os.Getenv("HOST")
}

// Porta para acesso ao servidor
func GetHTTPPort() string {
	return os.Getenv("HTTP_PORT")
}

// DNS para oferecer URL para servidor
func GetDNS() string {
	return os.Getenv("DNS")
}

// ValidateServerEnv valida todos os itens necessarios para iniciar o servidor
func ValidateServerEnv() error {
	if GetHostServer() == "" {
		return fmt.Errorf("variavel de ambiente HOST nao defindo no .env")
	}
	if GetHTTPPort() == "" {
		return fmt.Errorf("variavel de ambiente HTTP_PORT nao defindo no .env")
	}
	if GetDNS() == "" {
		return fmt.Errorf("variavel de ambiente DNS nao defindo no .env")
	}
	return nil
}
