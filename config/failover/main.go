package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strings"
)

const EtcdPath = "/appconfig/master"

var AppConfig *Config
var Conn *sql.DB

type Config struct {
	MysqlHost string
	MysqlPort string
	MysqlUser string
	MysqlPass string
}

func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(AppConfig)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func MysqlHealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if err := Conn.Ping(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to connect mysql"))
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func OpenSqlConnection() {
	var err error
	Conn, err = sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/",
			AppConfig.MysqlUser,
			AppConfig.MysqlPass,
			AppConfig.MysqlHost,
			AppConfig.MysqlPort,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func ReloadConfig(key string, val string) {
	key = strings.TrimPrefix(key, EtcdPath)
	switch key {
	case "/host":
		AppConfig.MysqlHost = val
	case "/port":
		AppConfig.MysqlPort = val
	case "/user":
		AppConfig.MysqlUser = val
	case "/pass":
		AppConfig.MysqlPass = val
	}

	OpenSqlConnection()
}

func WatchConfig(client *etcd.Client) {
	watchChan := make(chan *etcd.Response)
	go client.Watch(EtcdPath, 0, true, watchChan, nil)
	go func() {
		for {
			select {
			case resp := <-watchChan:
				log.Printf("Config is changed, %s, %s", resp.Node.Key, resp.Node.Value)
				ReloadConfig(resp.Node.Key, resp.Node.Value)
			}
		}
	}()
}

func main() {
	// Default Config
	AppConfig = &Config{
		MysqlHost: "192.168.99.100",
		MysqlPort: "3306",
		MysqlUser: "admin",
		MysqlPass: "admin",
	}

	OpenSqlConnection()

	client := etcd.NewClient([]string{"http://192.168.99.100:4001"})
	WatchConfig(client)

	http.HandleFunc("/", ConfigHandler)
	http.HandleFunc("/health", MysqlHealthCheckHandler)
	http.ListenAndServe(":3000", nil)
}
