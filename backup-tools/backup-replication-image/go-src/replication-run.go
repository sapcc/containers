package main

import (
	"os"
	"time"

	"github.com/sapcc/containers/backup-tools/go-src/prometheus"
)

var (
	// PromGauge is the prometheus pointer to use in the other files on same directory path
	PromGauge   *prometheus.Gauge
	DebugOutput = os.Getenv("DEBUG") == "yes"
)

func main() {
	PromGauge = prometheus.NewReplication()

	go func() {
		for {
			PromGauge.Beginn()
			LoadAndStartJobs()
			PromGauge.Finish()
			time.Sleep(7200 * time.Second)
		}
	}()

	PromGauge.ServerStart()
}
