package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"net/http"
)

func main() {
	// Get a new client, with KV endpoints
	client, _ := api.NewClient(&api.Config{
		Address:    "192.168.99.100:8500",
		Scheme:     "http",
		HttpClient: http.DefaultClient,
	})

	kv := client.KV()

	// Add new key/value
	p := &api.KVPair{Key: "foo", Value: []byte("bar")}
	_, err := kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Get value from foo
	pair, _, err := kv.Get("foo", nil)
	if err != nil {
		panic(err)
	}

	// Result is `bar`
	fmt.Printf(string(pair.Value))
}
