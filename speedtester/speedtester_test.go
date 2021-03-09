package speedtester

import "testing"

var download float64
var upload float64

func BenchmarkSpeedtestCom(b *testing.B) {
	for n := 0; n < b.N; n++ {
		download, upload, _ = GetSpeeds(SpeedTestCom)
	}
}

func BenchmarkFastCom(b *testing.B) {
	for n := 0; n < b.N; n++ {
		download, upload, _ = GetSpeeds(FastCom)
	}
}

func TestSpeedTestCreatingSpeettestComPositiveFlow(t *testing.T) {
	speedtestComTester, err := NewSpeedTester(SpeedTestCom)
	if err != nil {
		t.Errorf("Want error to be: nil\ngot %v\n", err)
	}

	if speedtestComTester == nil {
		t.Error("Want not nil speedTester got nil\n")
	}
}

func TestSpeedTestCreatingFastComPositiveFlow(t *testing.T) {
	fastComTester, err := NewSpeedTester(FastCom)
	if err != nil {
		t.Errorf("Want error to be: nil\ngot %v\n", err)
	}

	if fastComTester == nil {
		t.Error("Want not nil speedTester got nil\n")
	}
}

func TestSpeedTestCreatingWrongProvider(t *testing.T) {
	_, err := NewSpeedTester("DummyProvider")
	if _, ok := err.(UknownProviderError); !ok {
		t.Errorf("Want error to be of type : UknownProviderError\ngot %T\n", err)
	} else {
		t.Log("Got error: ", err.Error())
	}

}

func runIntegrationPositiveFlow(provider string, t *testing.T) {
	tester, err := NewSpeedTester(provider)

	if err != nil {
		t.Errorf("Want error to be: nil\ngot %v\n", err)
	}

	err = tester.Run()

	if err != nil {
		t.Errorf("Want error to be: nil\ngot %v\n", err)
	}

	if !(tester.GetDownloadSpeed() > 0.0) {
		t.Errorf("Want download speed to be: bigger than 0.0 \ngot %v\n", tester.GetDownloadSpeed())
	}

	if !(tester.GetUploadSpeed() > 0.0) {
		t.Errorf("Want upload speed to be: bigger than 0.0 \ngot %v\n", tester.GetUploadSpeed())
	}
}

func TestIntegrationSpeedtestComPositiveFlow(t *testing.T) {
	runIntegrationPositiveFlow(SpeedTestCom, t)
}

func TestIntegrationFastComPositiveFlow(t *testing.T) {
	runIntegrationPositiveFlow(FastCom, t)
}

func runIntegrationGetSpeedsPositiveFlow(provider string, t *testing.T) {
	download, upload, err := GetSpeeds(provider)

	if err != nil {
		t.Errorf("Want error to be: nil\ngot %v\n", err)
	}

	if !(download > 0.0) {
		t.Errorf("Want download speed to be: bigger than 0.0 \ngot %v\n", download)
	}

	if !(upload > 0.0) {
		t.Errorf("Want upload speed to be: bigger than 0.0 \ngot %v\n", upload)
	}
}

func TestIntegrationGetSpeedsSpeedtestComPositiveFlow(t *testing.T) {
	runIntegrationGetSpeedsPositiveFlow(SpeedTestCom, t)
}

func TestIntegrationGetSpeedsFastComPositiveFlow(t *testing.T) {
	runIntegrationGetSpeedsPositiveFlow(FastCom, t)
}
