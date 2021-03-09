package speedtester

import "fmt"

type SpeedTester interface {
	Run() error
	GetUploadSpeed() float64
	GetDownloadSpeed() float64
}

type ConfigurationError string

func (c ConfigurationError) Error() string {
	return string(c)
}

type DownloadTestError string

func (d DownloadTestError) Error() string {
	return string(d)
}

type UploadTestError string

func (u UploadTestError) Error() string {
	return string(u)
}

type UknownProviderError string

func (u UknownProviderError) Error() string {
	return fmt.Sprintf("Uknown provider: %s", u)
}
