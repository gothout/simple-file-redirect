package dto

import (
	"path/filepath"
	"strings"

	"simple-file-redirect/internal/app/model"
)

type UploadFileResponseDTO struct {
	File string `form:"file" binding:"required"`
}

type UploadaFileConvertResponseDTO struct {
	File    string `form:"file" binding:"required"`
	Convert string `form:"convert" binding:"required"`
}

// ToModel converte o DTO em um model.File, separando nome e extensão
func (dto *UploadFileResponseDTO) ToModel() model.File {
	ext := filepath.Ext(dto.File)
	name := strings.TrimSuffix(dto.File, ext)

	return model.File{
		Name: name,
		Ext:  strings.TrimPrefix(ext, "."), // tira o ponto
		Path: "",                           // path pode ser preenchido depois
	}
}

func (dto *UploadaFileConvertResponseDTO) ToModelConvert() model.FileConvert {
	ext := filepath.Ext(dto.File)
	name := strings.TrimSuffix(dto.File, ext)

	return model.FileConvert{
		Name:   name,
		ExtOr:  strings.TrimPrefix(ext, "."), // extensão original
		ExtDst: strings.ToLower(dto.Convert), // extensão destino
		Path:   "",                           // preenchido depois se necessário
	}
}
