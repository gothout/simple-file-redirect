package manager

import "mime/multipart"

type Service interface {
	SaveFile(fileHeader *multipart.FileHeader) (string, error)
	DeleteFile(filePath string) (string, error)
}
