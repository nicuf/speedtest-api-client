package speedtester

const (
	// SpeedTestCom is provider string for speedtest.com
	SpeedTestCom string = "speedtest.com"
	// FastCom is provider string for fast.com
	FastCom string = "fast.com"
)

// NewSpeedTester creates a SpeedTester for the provided test provider.
// It returns the created SpeedTester.
func NewSpeedTester(provider string) (SpeedTester, error) {
	if provider == SpeedTestCom {
		return newSpeedtestTester(), nil
	}

	if provider == FastCom {
		return newFastComtester(), nil
	}

	return nil, UknownProviderError(provider)
}

// GetSpeeds is a utility function to get directly the download and upload speed.
// It return the download and upload speed in Mbps, and an error.
func GetSpeeds(provider string) (float64, float64, error) {
	tester, err := NewSpeedTester(provider)
	if err != nil {
		return 0.0, 0.0, err
	}

	err = tester.Run()
	if err != nil {
		return 0.0, 0.0, err
	}

	return tester.GetDownloadSpeed(), tester.GetUploadSpeed(), nil
}
