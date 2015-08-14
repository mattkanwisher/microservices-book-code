// main.go
package main

import (
	"github.com/coreos/go-etcd/etcd"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

const EtcdPath = "/appconfig"

func main() {
	// Default Config
	AppConfig = &Config{
		Master: MysqlConfig{
			Host: "192.168.99.100",
			Port: "3306",
			User: "admin",
			Pass: "admin",
		},
		Slaves: []MysqlConfig{
			{
				Host: "192.168.99.100",
				Port: "3307",
				User: "admin",
				Pass: "admin",
			},
		},
	}

	OpenSqlConnection()

	CreateSchema()

	client := etcd.NewClient([]string{"http://192.168.99.100:4001"})
	WatchConfig(client)

	http.HandleFunc("/", ConfigHandler)
	http.HandleFunc("/health", MysqlHealthCheckHandler)
	http.HandleFunc("/genuser", GenUserHandler)
	http.HandleFunc("/user", ReadUserHandler)
	http.ListenAndServe(":3000", nil)
}
