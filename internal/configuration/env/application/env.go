package application

import (
	"fmt"
	"os"
)

// Token para as rotas da aplicação
func GetTokenApp() string {
	return os.Getenv("TOKEN_APPLICATION")
}

// ValidateApplicationEnv valida todos os itens necessarios para iniciar o servidor
func ValidateApplicationEnv() error {
	if GetTokenApp() == "" {
		return fmt.Errorf("variavel de ambiente TOKEN_APPLICATION nao defindo no .env")
	}
	return nil
}
