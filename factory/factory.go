package factory

type TestUtilityWrapper interface {
	Run() error
	GetUploadSpeed() float64
	GetDownloadSpeed() float64
}

type ConfigurationError string

func (c ConfigurationError) Error() string {
	return string(c)
}
