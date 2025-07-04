package converter

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) ConvertMP3toOGG(inputPath string) (string, error) {
	// Verifica se é um .mp3
	if !strings.HasSuffix(strings.ToLower(inputPath), ".mp3") {
		return "", fmt.Errorf("formato inválido: somente arquivos .mp3 são suportados")
	}

	// Gera caminho de saída com extensão .ogg
	outputPath := strings.TrimSuffix(inputPath, filepath.Ext(inputPath)) + ".ogg"

	// Comando para converter: ffmpeg -i input.mp3 output.ogg
	cmd := exec.Command("ffmpeg", "-y", "-i", inputPath, outputPath)

	// Executa o comando
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("erro ao converter arquivo: %v", err)
	}

	return outputPath, nil
}
