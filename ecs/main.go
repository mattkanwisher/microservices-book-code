package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

const (
	VERSION = "1"
)

var (
	appId       = uuid.NewV1().String()
	serviceName = os.Getenv("SERVICE_NAME")
	talkTo      = os.Getenv("TALK_TO")
	addr        = os.Getenv("ADDR")
	hostname    = os.Getenv("HOSTNAME")
)

func randomTalk() string {
	talks := strings.Split(talkTo, ",")
	return talks[rand.Int()%len(talks)]
}

func WhoAreYou(w http.ResponseWriter, r *http.Request) {
	resp, err := http.DefaultClient.Get(randomTalk() + "/answer")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	answer, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(answer)
}

func Answer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	str := fmt.Sprintf("I am %s version %s from app id %s $HOSTNAME %s",
		serviceName,
		VERSION,
		appId,
		hostname,
	)
	w.Write([]byte(str))
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func Env(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strings.Join(os.Environ(), "\n")))
}

func main() {
	http.HandleFunc("/whoareyou", WhoAreYou)
	http.HandleFunc("/answer", Answer)
	http.HandleFunc("/health", Health)
	http.HandleFunc("/env", Env)

	if addr == "" {
		addr = ":8080"
	}

	http.ListenAndServe(addr, nil)
}
