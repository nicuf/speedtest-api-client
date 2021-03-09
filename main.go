package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.jonnrb.io/speedtest/fastdotcom"
	"go.jonnrb.io/speedtest/oututil"
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
		fmt.Printf("Download: %f, Upload: %f\n", s.DLSpeed, s.ULSpeed)
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

	formatFunc := func(s units.BytesPerSecond) string {
		return formatSpeed("Download speed", s)
	}

	stream, finalize := proberPrinter(formatFunc)

	speed, err := m.ProbeDownloadSpeed(ctx, client, stream)
	if err != nil {
		log.Fatalf("Error probing download speed: %v", err)
		return
	}
	finalize(speed)

	fmt.Println("Final Download Speed: ", speed)
}

func upload(m *fastdotcom.Manifest, client *fastdotcom.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	formatFunc := func(s units.BytesPerSecond) string {
		return formatSpeed("Upload speed", s)
	}

	stream, finalize := proberPrinter(formatFunc)

	speed, err := m.ProbeUploadSpeed(ctx, client, stream)
	if err != nil {
		log.Fatalf("Error probing upload speed: %v", err)
		return
	}
	finalize(speed)

	fmt.Println("Final Upload Speed: ", speed)
}

func proberPrinter(format func(units.BytesPerSecond) string) (
	stream chan units.BytesPerSecond,
	finalize func(units.BytesPerSecond),
) {
	p := oututil.StartPrinting()
	p.Println(format(units.BytesPerSecond(0)))

	stream = make(chan units.BytesPerSecond)
	var g errgroup.Group
	g.Go(func() error {
		for speed := range stream {
			p.Println(format(speed))
		}
		return nil
	})

	finalize = func(s units.BytesPerSecond) {
		g.Wait()
		p.Finalize(format(s))
	}
	return
}

func formatSpeed(prefix string, s units.BytesPerSecond) string {
	i := s.BitsPerSecond()
	return fmt.Sprintf("%s: %v", prefix, i)
}
