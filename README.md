# speedtest-api-client

This module aim is to get download and upload speeds by using Ookla's https://www.speedtest.net/ and Netflix's https://fast.com/

## Usage 
It can be used with the help of utility function:
```go
download, upload, err := GetSpeeds(provider)
```
where provider is either "speedtest.com" or "fast.com".

or if more control over the returned errors is needed a speed tester object can be used:

```go
tester, err := NewSpeedTester(provider)
if err != nil {
    fmt.Println("Got an error durring creation of a speed tester: ", err)
}

err = tester.Run()
if err != nil {
    fmt.Println("Got an error when runing the speed tests: ", err)
}

downloadSpeed := tester.GetDownloadSpeed()
uploadSpeed := tester.GetUploadSpeed()
```

Full documentation can be can be found after running ``` docker-compose up``` on [Speedtest-Api-Client](http://localhost:6060/pkg/github.com/nicuf/speedtest-api-client/speedtester/) .