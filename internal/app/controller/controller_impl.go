package controller

import (
	"net/http"
	"path/filepath"
	"strings"

	"simple-file-redirect/internal/app/binding"
	"simple-file-redirect/internal/app/service"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service service.Service
}

func NewController(service service.Service) Controller {
	return &controller{
		service: service,
	}
}

// UploadArquivo godoc
// @Summary      Upload de arquivo
// @Description  Recebe um arquivo via multipart/form e salva no diretório de arquivos
// @Tags         Arquivos
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        file formData file true "Arquivo para upload"
// @Success      200 {object} map[string]string "Upload realizado com sucesso"
// @Failure      400 {object} map[string]string "Erro de validação"
// @Failure      500 {object} map[string]string "Erro ao salvar"
// @Router       /manager/v1/upload [post]
func (ctrl *controller) UploadArquivo(c *gin.Context) {
	_, fileHeader, err := binding.BindUploadFileDTO(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Arquivo e/ou campos obrigatórios não enviados"})
		return
	}

	fileModel, err := ctrl.service.SaveFile(fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao salvar o arquivo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "Upload realizado com sucesso",
		"name": fileModel.Name,
		"ext":  fileModel.Ext,
		"path": fileModel.Path,
	})
}

// DownloadArquivo godoc
// @Summary      Download de arquivo
// @Description  Realiza o download de um arquivo salvo, baseado no path informado
// @Tags         Arquivos
// @Produce      octet-stream
// @Param        path query string true "Caminho completo do arquivo salvo"
// @Param				 token query string true "Token para download do arquivo salvo"
// @Success      200 {file} file "Arquivo enviado"
// @Failure      400 {object} map[string]string "Parâmetro ausente"
// @Failure      404 {object} map[string]string "Arquivo não encontrado"
// @Router       /manager/v1/download [get]
func (ctrl *controller) DownloadArquivo(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Parâmetro 'path' é obrigatório"})
		return
	}

	normalized := filepath.ToSlash(path)

	for strings.Contains(normalized, "//") {
		normalized = strings.ReplaceAll(normalized, "//", "/")
	}

	if !strings.HasPrefix(normalized, "internal/storage/files/") {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Caminho inválido: deve estar dentro de 'internal/storage/files/'"})
		return
	}

	file, err := ctrl.service.DownloadFile(path)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Arquivo não encontrado"})
		return
	}
	defer file.Close()

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename=\""+filepath.Base(path)+"\"")
	c.Header("Content-Type", "application/octet-stream")

	c.File(path)

	go ctrl.service.DeleteFile(path)
}

// ListenArquivo godoc
// @Summary      Ouvir arquivo
// @Description  Retorna o arquivo de áudio para ser reproduzido diretamente, com possibilidade de remoção automática após execução
// @Tags         Arquivos
// @Produce      audio/mpeg
// @Produce      audio/ogg
// @Produce      application/octet-stream
// @Param        path query string true "Caminho completo do arquivo salvo"
// @Param        token query string true "Token para acesso ao arquivo"
// @Success      200 {file} file "Arquivo de áudio retornado"
// @Failure      400 {object} map[string]string "Parâmetro ausente ou inválido"
// @Failure      404 {object} map[string]string "Arquivo não encontrado"
// @Router       /manager/v1/listen [get]
func (ctrl *controller) ListenArquivo(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Parâmetro 'path' é obrigatório"})
		return
	}

	file, err := ctrl.service.DownloadFile(path)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Arquivo não encontrado"})
		return
	}
	defer file.Close()

	c.File(path)
	go ctrl.service.DeleteFile(path)
}

// ConvertArquivo godoc
// @Summary      Conversão de arquivo MP3 para OGG
// @Description  Realiza upload e conversão de um arquivo MP3 para OGG
// @Tags         Conversão
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        file formData file true "Arquivo MP3 para conversão"
// @Param        convert formData string true "Extensão de destino (ex: ogg)"
// @Success      200 {object} map[string]string "Arquivo convertido com sucesso"
// @Failure      400 {object} map[string]string "Erro de validação ou tipo de conversão não suportado"
// @Failure      500 {object} map[string]string "Erro interno ao converter"
// @Router       /manager/v1/convert [post]
func (ctrl *controller) ConvertArquivo(c *gin.Context) {
	dtoConv, fileHeader, err := binding.BindUploadFileConvertDTO(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Arquivo e/ou campo 'convert' obrigatórios"})
		return
	}

	fileModel, err := ctrl.service.SaveFile(fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao salvar o arquivo"})
		return
	}

	modelConv := dtoConv.ToModelConvert()
	//modelConv.Path = fileModel.Path

	if strings.ToLower(modelConv.ExtOr) != "mp3" || strings.ToLower(modelConv.ExtDst) != "ogg" {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Conversão não suportada. Apenas mp3 para ogg é permitido"})
		return
	}

	convertedPath, err := ctrl.service.ConvertMP3toOGG(fileModel.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao converter arquivo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":            "Arquivo convertido com sucesso",
		"original_name":  fileModel.Name,
		"original_ext":   fileModel.Ext,
		"converted_path": convertedPath,
	})
}
