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
