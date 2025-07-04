package server

import (
	"crypto/tls"
	"fmt"
	"os"
	"strings"
)

// Host para o servidor
func GetHostServer() string {
	return os.Getenv("HOST")
}

// GetHTTPSuse checks if HTTPS should be used and if certificates exist and are valid.
func GetHTTPSuse() bool {
	https := strings.ToUpper(os.Getenv("HTTPS"))
	if https != "TRUE" {
		return false
	}

	certPath := "./certificates/cert.crt"
	keyPath := "./certificates/privkey.key"

	// Check if both certificate files exist
	if _, err := os.Stat(certPath); os.IsNotExist(err) {
		fmt.Println("Certificate file not found:", certPath)
		return false
	}
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		fmt.Println("Private key file not found:", keyPath)
		return false
	}

	// Try loading the certificate to validate it
	_, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		fmt.Println("Invalid TLS certificate or key:", err)
		return false
	}

	return true
}

// Porta para acesso ao servidor HTTP
func GetHTTPPort() string {
	return os.Getenv("HTTP_PORT")
}

// GetHTTPSPort retrieves the HTTPS_PORT environment variable.
func GetHTTPSPort() string {
	return os.Getenv("HTTPS_PORT")
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
	if GetHTTPSPort() == "" {
		return fmt.Errorf("variavel de ambiente HTTPS_PORT nao defindo no .env")
	}
	if GetDNS() == "" {
		return fmt.Errorf("variavel de ambiente DNS nao defindo no .env")
	}
	https := strings.ToUpper(os.Getenv("HTTPS"))
	if https != "TRUE" && https != "FALSE" {
		return fmt.Errorf("variavel de ambiente HTTPS deve ser TRUE ou FALSE")
	}
	if https == "TRUE" && !GetHTTPSuse() {
		return fmt.Errorf("HTTPS esta ativo porem o certificado nao e valido")
	}
	return nil
}
