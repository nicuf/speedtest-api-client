package speedtester

import (
	"github.com/showwin/speedtest-go/speedtest"
)

type speedTestCom struct {
	downloadSpeed float64
	uploadSpeed   float64
}

func NewSpeedtestTester() SpeedTester {
	return &speedTestCom{}
}

func (stc *speedTestCom) Run() error {

	user, err := speedtest.FetchUserInfo()

	if err != nil {
		return ConfigurationError(err.Error())
	}

	serverList, err := speedtest.FetchServerList(user)

	if err != nil {
		return ConfigurationError(err.Error())
	}

	targets, err := serverList.FindServer([]int{})

	if err != nil {
		return ConfigurationError(err.Error())
	}

	for _, s := range targets {
		err = s.DownloadTest(false)
		if err != nil {
			return DownloadTestError(err.Error())
		}

		err = s.UploadTest(false)
		if err != nil {
			return UploadTestError(err.Error())
		}

		stc.downloadSpeed = s.DLSpeed
		stc.uploadSpeed = s.ULSpeed
	}

	return nil
}

func (stc *speedTestCom) GetDownloadSpeed() float64 {
	return stc.downloadSpeed
}

func (stc *speedTestCom) GetUploadSpeed() float64 {
	return stc.uploadSpeed
}
