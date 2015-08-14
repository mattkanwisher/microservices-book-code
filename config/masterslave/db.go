// db.go
package main

import (
	"fmt"
	"github.com/tsenart/nap"
	"log"
	"strings"
)

var Conn *nap.DB

func CreateSchema() {
	query := `CREATE DATABASE IF NOT EXISTS user`
	if _, err := Conn.Exec(query); err != nil {
		panic(err)
	}

	query = `CREATE TABLE IF NOT EXISTS user.user (
        id varchar(10),
        user varchar(100),
        PRIMARY KEY (id)
    );`
	if _, err := Conn.Exec(query); err != nil {
		panic(err)
	}
}

func OpenSqlConnection() {
	var err error
	dnsFormat := "%s:%s@tcp(%s:%s)/"

	dnsList := []string{
		fmt.Sprintf(dnsFormat,
			AppConfig.Master.User,
			AppConfig.Master.Pass,
			AppConfig.Master.Host,
			AppConfig.Master.Port,
		),
	}

	for i := 0; i < len(AppConfig.Slaves); i++ {
		dnsList = append(dnsList, fmt.Sprintf(dnsFormat,
			AppConfig.Slaves[i].User,
			AppConfig.Slaves[i].Pass,
			AppConfig.Slaves[i].Host,
			AppConfig.Slaves[i].Port,
		))
	}

	Conn, err = nap.Open("mysql", strings.Join(dnsList, ";"))
	if err != nil {
		log.Fatal(err)
	}
}
