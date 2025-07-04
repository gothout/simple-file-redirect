package converter

type Service interface {
	ConvertMP3toOGG(inputPath string) (string, error)
}
