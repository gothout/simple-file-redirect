package manager

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

const dirPath = "./internal/storage/files"

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) SaveFile(fileHeader *multipart.FileHeader) (string, error) {
	// cria pasta caso nao exista
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("erro ao criar diretorio: %w", err)
	}

	// abre o arquivo enviado
	srcFile, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("erro ao abrir arquivo: %w", err)
	}
	defer srcFile.Close()

	// define o caminho completo onde o arquivo sera salvo
	dstPath := filepath.Join(dirPath, fileHeader.Filename)
	// cria o arquivo de destino
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return "", fmt.Errorf("erro ao criar arquivo: %w", err)
	}
	defer dstFile.Close()

	// copia o conteudo
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return "", fmt.Errorf("erro ao salvar o arquivo: %v", err)
	}

	// Retorna o caminho salvo
	return dstPath, nil

}

func (s *service) DeleteFile(filePath string) (bool, error) {
	// Verifica se o arquivo existe
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false, fmt.Errorf("arquivo não encontrado: %s", filePath)
	}

	// Remove o arquivo
	if err := os.Remove(filePath); err != nil {
		return false, fmt.Errorf("erro ao deletar arquivo: %w", err)
	}

	return true, nil
}
