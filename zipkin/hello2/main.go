package main

import (
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"gopkg.in/spacemonkeygo/monitor.v1/trace"
	"gopkg.in/spacemonkeygo/monitor.v1/trace/gen-go/zipkin"
)

func HelloHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	RandomSQLCall(ctx)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world #2"))
}

func RandomSQLCall(ctx context.Context) {
	defer trace.Trace(&ctx)(nil)
	time.Sleep(time.Duration(rand.Int()%10000) * time.Microsecond)
}

func main() {
	c, err := trace.NewUDPCollector("127.0.0.1:3100", 128)
	if err != nil {
		panic(err)
	}
	trace.Configure(1, true, &zipkin.Endpoint{
		Ipv4:        127*256*256*256 + 1,
		Port:        3002,
		ServiceName: "hello2"})
	trace.RegisterTraceCollector(c)

	handler := trace.ContextWrapper(trace.TraceHandler(trace.ContextHTTPHandlerFunc(HelloHandler)))

	http.ListenAndServe(":3002", handler)
}
