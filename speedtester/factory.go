package speedtester

const (
	SpeedTestCom string = "speedtest.com"
	FastCom      string = "fast.com"
)

func NewSpeedTester(provider string) (SpeedTester, error) {
	if provider == SpeedTestCom {
		return NewSpeedtestTester(), nil
	}

	if provider == FastCom {
		return NewFastComtester(), nil
	}

	return nil, UknownProviderError(provider)
}
