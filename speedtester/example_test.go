package speedtester

import "fmt"

func ExampleGetSpeeds() {
	download, upload, err := GetSpeeds(SpeedTestCom)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("speedtest.com results: download: %.2f Mbps, upload: %.2f Mbps\n", download, upload)

	download, upload, err = GetSpeeds(FastCom)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("fast.com results: download: %.2f Mbps, upload: %.2f Mbps\n", download, upload)
}

func ExampleNewSpeedTester() {
	tester, err := NewSpeedTester(FastCom)
	if err != nil {
		fmt.Println("Got an error durring creation of a speed tester: ", err)
	}

	err = tester.Run()
	if err != nil {
		fmt.Println("Got an error when runing the speed tests: ", err)
	}

	fmt.Printf("Download: %.2fMbps, Upload: %.2fMbps\n", tester.GetDownloadSpeed(), tester.GetUploadSpeed())
}
