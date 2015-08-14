package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/fzzy/radix/redis"
)

type XKCD struct {
	Number   int    `json:"num"`
	Alt      string `json:"alt"`
	ImageURL string `json:"img"`
}

const XKCDLatestURL = "http://xkcd.com/info.0.json"
const XKCDURLFormat = "http://xkcd.com/%d/info.0.json"

var client *redis.Client

func main() {
	var err error
	rand.Seed(time.Now().Unix())

	client, err = redis.Dial("tcp", "0.0.0.0:6379")
	noError(err)
	reply, err := client.Cmd("PING").Str()
	noError(err)
	if reply != "PONG" {
		panic("Could not connect to Redis!")
	}

	latest := &XKCD{}
	noError(get(XKCDLatestURL, latest))

	number := rand.Int31n(int32(latest.Number)) + 1
	comic, url := &XKCD{}, fmt.Sprintf(XKCDURLFormat, number)
	noError(get(url, comic))

	fmt.Println(comic.Number, comic.Alt)
	fmt.Println(comic.ImageURL)
}

func get(url string, result interface{}) error {
	if reply := client.Cmd("GET", url); reply.Err != nil {
		return reply.Err
	} else if reply.Type != redis.NilReply {
		buffer, _ := reply.Bytes()
		return json.Unmarshal(buffer, result)
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bytes, result); e != nil {
		return err
	}

	if reply, err := client.Cmd("SET", url, bytes, "EX", 24*60*60).Str(); e != nil {
		return err
	} else if reply != "OK" {
		return fmt.Errorf("failed to cache a url.")
	}

	return nil
}

func noError(err error) {
	if err != nil {
		panic(err)
	}
}
