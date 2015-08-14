package main

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
)

func main() {
	machines := []string{"http://192.168.99.100:4001"}
	client := etcd.NewClient(machines)

	// Add new key/value
	if _, err := client.Set("/message", "Hello world", 0); err != nil {
		panic(err)
	}

	// Get value from /foo
	resp, err := client.Get("/message", false, false)
	if err != nil {
		panic(err)
	}

	// Result is `bar`
	fmt.Println(resp.Node.Value)
}
