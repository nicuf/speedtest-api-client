package speedtester

// SpeedTester is the interface for to interact with different types of speed testers
type SpeedTester interface {
	Run() error
	GetUploadSpeed() float64
	GetDownloadSpeed() float64
}

// ConfigurationError is returned when there is an error during speed testers configuration phase.
type ConfigurationError string

func (c ConfigurationError) Error() string {
	return string(c)
}

// DownloadTestError is returned when there is an error during testing download speed.
type DownloadTestError string

func (d DownloadTestError) Error() string {
	return string(d)
}

// UploadTestError is returned when there is an error during testing upload speed.
type UploadTestError string

func (u UploadTestError) Error() string {
	return string(u)
}

// UknownProviderError is returned when theere is no speed tester for the provided provider.
type UknownProviderError string

func (u UknownProviderError) Error() string {
	return string("Uknown provider: " + u)
}
