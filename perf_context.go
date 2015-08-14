package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/context"
)

func contextHandler(w http.ResponseWriter, req *http.Request) {
	var timeInMilliseconds time.Duration = 0
	ctx := context.WithValue(context.Background(), "time", &timeInMilliseconds)

	longTimeSqlFunc(ctx)
	longTimeSqlFunc(ctx)
	longTimeSqlFunc(ctx)

	val := ctx.Value("time").(*time.Duration)
	s := fmt.Sprintf("Took in expensive functions(%f)\n", val.Seconds()/1000)
	w.Write([]byte(s))
}

func longTimeSqlFunc(ctx context.Context) {
	defer func(s time.Time) { 
		val := ctx.Value("time").(*time.Duration)
		*val = *val + time.Since(s)
	}(time.Now())

	time.Sleep(time.Second)
}

func main() {
    http.HandleFunc("/", contextHandler)
    http.ListenAndServe(":8080", nil)
}