package service

import (
	"mime/multipart"
	"os"
	"simple-file-redirect/internal/app/model"
)

type Service interface {
	SaveFile(fileHeader *multipart.FileHeader) (*model.File, error)
	DownloadFile(path string) (*os.File, error)
	ConvertMP3toOGG(path string) (string, error)
}
