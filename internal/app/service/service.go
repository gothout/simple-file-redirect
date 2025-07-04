package service

import (
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"simple-file-redirect/internal/app/model"
	"simple-file-redirect/internal/storage/converter"
	"simple-file-redirect/internal/storage/manager"
)

type service struct {
	storageManager   manager.Service
	storageConverter converter.Service
}

func NewService(storageManager manager.Service, storageConverter converter.Service) Service {
	return &service{
		storageManager:   storageManager,
		storageConverter: storageConverter,
	}
}

func (s *service) SaveFile(fileHeader *multipart.FileHeader) (*model.File, error) {
	// Salva o arquivo e recebe o caminho real
	path, err := s.storageManager.SaveFile(fileHeader)
	if err != nil {
		return nil, err
	}

	// Extrai nome e extensão
	ext := strings.TrimPrefix(filepath.Ext(fileHeader.Filename), ".")
	name := strings.TrimSuffix(fileHeader.Filename, filepath.Ext(fileHeader.Filename))

	// Cria o model com base nas informações
	fileModel := &model.File{
		Name: name,
		Ext:  ext,
		Path: path,
	}

	return fileModel, nil
}

func (s *service) DownloadFile(path string) (*os.File, error) {
	// Verifica se o arquivo existe
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	// Abre o arquivo
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *service) ConvertMP3toOGG(path string) (string, error) {
	resultPath, err := s.storageConverter.ConvertMP3toOGG(path)
	if err != nil {
		return "", err
	}

	// Deleta o original
	_, _ = s.storageManager.DeleteFile(path)

	return resultPath, nil
}
