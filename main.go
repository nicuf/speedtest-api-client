package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.jonnrb.io/speedtest/fastdotcom"
	"go.jonnrb.io/speedtest/units"
	"golang.org/x/sync/errgroup"

	spdtest "github.com/showwin/speedtest-go/speedtest"
)

/*
type SpeedTester interface {
	GetSpeeds(ctx context.Context) (float64, float64)
}

type TestUtilityWrapper interface {
	Run(ctx context.Context) error
	GetUploadSpeed() float64
	GetDownloadSpeed() float64
}

type speedTestWrapper struct {
}

*/

func main() {

	user, _ := spdtest.FetchUserInfo()

	serverList, _ := spdtest.FetchServerList(user)
	targets, _ := serverList.FindServer([]int{})

	for _, s := range targets {
		s.DownloadTest(false)
		s.UploadTest(false)
		fmt.Printf("Download: %f Mb/s, Upload: %fMb/s\n", s.DLSpeed, s.ULSpeed)
	}

	var client fastdotcom.Client

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	m, err := fastdotcom.GetManifest(ctx, 5)
	if err != nil {
		log.Fatalf("Error loading fast.com configuration: %v", err)
	}

	download(m, &client)
	upload(m, &client)
}

func download(m *fastdotcom.Manifest, client *fastdotcom.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, finalize := proberPrinter()

	speed, err := m.ProbeDownloadSpeed(ctx, client, stream)
	if err != nil {
		log.Fatalf("Error probing download speed: %v", err)
		return
	}
	finalize()

	fmt.Println("Final Download Speed: ", speed)
}

func upload(m *fastdotcom.Manifest, client *fastdotcom.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, finalize := proberPrinter()

	speed, err := m.ProbeUploadSpeed(ctx, client, stream)
	if err != nil {
		log.Fatalf("Error probing upload speed: %v", err)
		return
	}
	finalize()

	fmt.Println("Final Upload Speed: ", speed)
}

func proberPrinter() (
	stream chan units.BytesPerSecond,
	finalize func(),
) {

	stream = make(chan units.BytesPerSecond)
	var g errgroup.Group
	g.Go(func() error {
		for range stream {
		}
		return nil
	})

	finalize = func() {
		g.Wait()
	}
	return
}
