package main

import (
	"bytes"
	"net/http"

	"golang.org/x/net/context"
	"gopkg.in/spacemonkeygo/monitor.v1/trace"
	"gopkg.in/spacemonkeygo/monitor.v1/trace/gen-go/zipkin"
)

func HelloHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest("GET", "http://localhost:3002", nil)

	resp, err := trace.TraceRequest(ctx, http.DefaultClient, req)
	if err != nil {
		panic(err)
	}

	buf := bytes.Buffer{}
	buf.ReadFrom(resp.Body)
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	buf.WriteTo(w)
}

func main() {
	c, err := trace.NewUDPCollector("127.0.0.1:3100", 128)
	if err != nil {
		panic(err)
	}

	trace.Configure(1, true, &zipkin.Endpoint{
		Ipv4:        127*256*256*256 + 1,
		Port:        3001,
		ServiceName: "hello1"})
	trace.RegisterTraceCollector(c)

	handler := trace.ContextWrapper(trace.TraceHandler(trace.ContextHTTPHandlerFunc(HelloHandler)))

	http.ListenAndServe(":3001", handler)
}
