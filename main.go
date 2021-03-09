package main

import (
	"fmt"

	"github.com/nicuf/speedtest-api-client/testers"
)

/*
type SpeedTester interface {
	GetSpeeds(ctx context.Context) (float64, float64)
}



type speedTestWrapper struct {
}

*/

func main() {
	speedTestCom := testers.NewSpeedtestTester()
	err := speedTestCom.Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("speedtest.com results: download: %.2f Mbps, upload: %.2f Mbps\n", speedTestCom.GetDownloadSpeed(), speedTestCom.GetUploadSpeed())

	fastCom := testers.NewFastComtester()
	err = fastCom.Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("fast.com results: download: %.2f Mbps, upload: %.2f Mbps\n", fastCom.GetDownloadSpeed(), fastCom.GetUploadSpeed())
}
