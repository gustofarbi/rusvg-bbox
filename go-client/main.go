package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"github.com/schollz/progressbar/v3"
	"go-client/internal/svgbbox"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math"
	"os"
	"sync"
	"time"
)

var (
	conc        = flag.Int("conc", 50, "number of concurrent requests")
	numRequests = flag.Int("n", 1000, "number of requests")
)

func main() {
	flag.Parse()

	url := "localhost:50051"
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := svgbbox.NewCalculateBBoxClient(conn)
	wg := sync.WaitGroup{}
	b, err := os.ReadFile("../testdata/heart.svg")
	if err != nil {
		panic(err)
	}

	times := make([]time.Duration, 0, *numRequests)
	pb := progressbar.New(*numRequests)

	pool, err := ants.NewPoolWithFunc(*conc, func(interface{}) {
		start := time.Now()
		defer wg.Done()

		request := svgbbox.Svg{}
		request.Content = string(b)

		bbox, err := client.GetBBox(context.Background(), &request)
		if err != nil {
			panic(err)
		}
		if bbox == nil {
			panic("bbox is nil")
		}

		if err = pb.Add(1); err != nil {
			panic(err)
		}
		times = append(times, time.Since(start))
	})

	if err != nil {
		panic(err)
	}

	start := time.Now()
	for i := 0; i < *numRequests; i++ {
		wg.Add(1)
		if err = pool.Invoke(i); err != nil {
			panic(err)
		}
	}

	wg.Wait()
	pool.Release()

	fmt.Println()
	fmt.Printf("\ntotal time: %s\n", time.Since(start))
	mean, maxDur, minDur, stddev := calculateStats(times)
	fmt.Printf("\nmean: %v\nmaxDur: %v\nminDur: %v\nstddev: %v\n", mean, maxDur, minDur, stddev)
}

func calculateStats(durations []time.Duration) (mean, max, min, stddev time.Duration) {
	if len(durations) == 0 {
		return 0, 0, 0, 0
	}

	// Calculate mean
	var sum time.Duration
	for _, duration := range durations {
		sum += duration
	}
	mean = time.Duration(int64(sum) / int64(len(durations)))

	// Calculate max and min
	max = durations[0]
	min = durations[0]
	for _, duration := range durations {
		if duration > max {
			max = duration
		}
		if duration < min {
			min = duration
		}
	}

	// Calculate standard deviation
	var squaredDiffSum int64
	for _, duration := range durations {
		diff := int64(duration - mean)
		squaredDiffSum += diff * diff
	}
	variance := time.Duration(squaredDiffSum / int64(len(durations)))
	stddev = time.Duration(int64(math.Sqrt(float64(variance))))

	return mean, max, min, stddev
}
