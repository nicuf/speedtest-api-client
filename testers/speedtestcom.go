package testers

import (
	"github.com/nicuf/speedtest-api-client/factory"
	"github.com/showwin/speedtest-go/speedtest"
)

type speedTestCom struct {
	downloadSpeed float64
	uploadSpeed   float64
}

func NewSpeedtestTester() factory.TestUtilityWrapper {
	return &speedTestCom{}
}

func (stc *speedTestCom) Run() error {

	user, err := speedtest.FetchUserInfo()

	if err != nil {
		return factory.ConfigurationError(err.Error())
	}

	serverList, err := speedtest.FetchServerList(user)

	if err != nil {
		return factory.ConfigurationError(err.Error())
	}

	targets, err := serverList.FindServer([]int{})

	if err != nil {
		return factory.ConfigurationError(err.Error())
	}

	for _, s := range targets {
		err = s.DownloadTest(false)
		err = s.UploadTest(false)
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
