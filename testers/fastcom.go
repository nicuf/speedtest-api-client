package testers

import (
	"context"
	"time"

	"github.com/nicuf/speedtest-api-client/factory"
	"go.jonnrb.io/speedtest/fastdotcom"
	"go.jonnrb.io/speedtest/units"
)

const (
	kDefaultUrlCount int     = 5
	Kbps             float64 = 1000
	Mbps                     = 1000 * Kbps
)

type fastComTester struct {
	downloadSpeed float64
	uploadSpeed   float64
}

func NewFastComtester() factory.TestUtilityWrapper {
	return &fastComTester{}
}

func (fc *fastComTester) Run() error {
	var client fastdotcom.Client

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	m, err := fastdotcom.GetManifest(ctx, kDefaultUrlCount)
	if err != nil {
		return factory.ConfigurationError(err.Error())
	}

	err = fc.testDownload(m, &client)
	if err != nil {
		return factory.ConfigurationError(err.Error())
	}

	err = fc.testUpload(m, &client)
	if err != nil {
		return factory.ConfigurationError(err.Error())
	}

	return nil
}

func (fc *fastComTester) GetDownloadSpeed() float64 {
	return fc.downloadSpeed
}

func (fc *fastComTester) GetUploadSpeed() float64 {
	return fc.uploadSpeed
}

func (fc *fastComTester) testDownload(m *fastdotcom.Manifest, client *fastdotcom.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream := make(chan units.BytesPerSecond)
	go func() {
		for range stream {
		}
	}()

	speed, err := m.ProbeDownloadSpeed(ctx, client, stream)
	if err != nil {
		return err
	}

	fc.downloadSpeed = float64(speed.BitsPerSecond()) / Mbps

	return nil
}

func (fc *fastComTester) testUpload(m *fastdotcom.Manifest, client *fastdotcom.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream := make(chan units.BytesPerSecond)
	go func() {
		for range stream {
		}
	}()

	speed, err := m.ProbeUploadSpeed(ctx, client, stream)
	if err != nil {
		return err
	}

	fc.uploadSpeed = float64(speed.BitsPerSecond()) / Mbps

	return nil
}
