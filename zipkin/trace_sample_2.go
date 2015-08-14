package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
	"gopkg.in/spacemonkeygo/monitor.v1/trace"
	"gopkg.in/spacemonkeygo/monitor.v1/trace/gen-go/zipkin"
)

func main() {
	trace.Configure(1, true, &zipkin.Endpoint{
		Ipv4:        127*0x01000000 + 1,
		Port:        8080,
		ServiceName: "go-zipkin-testclient",
	})

	if c, e := trace.NewScribeCollector("0:9410"); e != nil {
		panic(e)
	} else {
		trace.RegisterTraceCollector(c)
	}

	fmt.Println(outer())
}

func outer() int {
	ctx := context.Background()
	defer trace.Trace(&ctx)(nil)

	time.Sleep(1 * time.Second)
	return inner1(ctx) + inner2(ctx)
}

func inner1(ctx context.Context) int {
	defer trace.Trace(&ctx)(nil)
	time.Sleep(1 * time.Second)
	return 1
}

func inner2(ctx context.Context) int {
	defer trace.Trace(&ctx)(nil)
	time.Sleep(1 * time.Second)
	return 1
}
