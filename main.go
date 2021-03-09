package main

import (
	"fmt"

	"github.com/nicuf/speedtest-api-client/speedtester"
)

/*
type SpeedTester interface {
	GetSpeeds(ctx context.Context) (float64, float64)
}



type speedTestWrapper struct {
}

*/

func main() {
	speedTestCom, err := speedtester.NewSpeedTester(speedtester.SpeedTestCom)
	if err != nil {
		fmt.Println(err)
	}

	err = speedTestCom.Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("speedtest.com results: download: %.2f Mbps, upload: %.2f Mbps\n", speedTestCom.GetDownloadSpeed(), speedTestCom.GetUploadSpeed())

	fastCom, err := speedtester.NewSpeedTester(speedtester.FastCom)
	if err != nil {
		fmt.Println(err)
	}

	err = fastCom.Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("fast.com results: download: %.2f Mbps, upload: %.2f Mbps\n", fastCom.GetDownloadSpeed(), fastCom.GetUploadSpeed())
}
