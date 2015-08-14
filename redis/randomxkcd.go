package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type XKCD struct {
	Number   int    `json:"num"`
	Alt      string `json:"alt"`
	ImageURL string `json:"img"`
}

const XKCDLatestURL = "http://xkcd.com/info.0.json"
const XKCDURLFormat = "http://xkcd.com/%d/info.0.json"

func main() {
	rand.Seed(time.Now().Unix())
	latest := &XKCD{}
	noError(get(XKCDLatestURL, latest))

	number := rand.Int31n(int32(latest.Number)) + 1
	comic, url := &XKCD{}, fmt.Sprintf(XKCDURLFormat, number)
	noError(get(url, comic))

	fmt.Println(comic.Number, comic.Alt)
	fmt.Println(comic.ImageURL)
}

func get(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(result)
}

func noError(err error) {
	if err != nil {
		panic(err)
	}
}
