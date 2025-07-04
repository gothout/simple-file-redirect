package binding

import (
	"fmt"
	"mime/multipart"
	"simple-file-redirect/internal/app/dto"

	"github.com/gin-gonic/gin"
)

func BindUploadFileDTO(c *gin.Context) (*dto.UploadFileResponseDTO, *multipart.FileHeader, error) {
	// Pega o arquivo direto do form
	file, err := c.FormFile("file")
	if err != nil {
		return nil, nil, err
	}

	// Preenche manualmente o DTO com o nome do arquivo
	dtoUpload := &dto.UploadFileResponseDTO{
		File: file.Filename,
	}

	return dtoUpload, file, nil
}

func BindUploadFileConvertDTO(c *gin.Context) (*dto.UploadaFileConvertResponseDTO, *multipart.FileHeader, error) {
	// Pega o arquivo
	file, err := c.FormFile("file")
	if err != nil {
		return nil, nil, err
	}

	// Pega o valor da conversão como string
	convert := c.PostForm("convert")
	if convert == "" {
		return nil, nil, fmt.Errorf("campo 'convert' é obrigatório")
	}

	// Preenche o DTO
	dto := &dto.UploadaFileConvertResponseDTO{
		File:    file.Filename,
		Convert: convert,
	}

	return dto, file, nil
}
