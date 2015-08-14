package main

import (
	// "github.com/spacemonkeygo/spacelog"
	"gopkg.in/spacemonkeygo/monitor.v1/trace"
)

func main() {
	// spacelog.MustSetup("collector", spacelog.SetupConfig{
	// 	Level:  "info",
	// 	Format: "{{ColorizeLevel .Level}}{{.Message}}{{.Reset}}"})

	collector, err := trace.NewScribeCollector("128.199.223.6:9410")
	if err != nil {
		panic(err)
	}
	defer collector.Close()

	panic(trace.RedirectPackets("127.0.0.1:3100", collector))
}
