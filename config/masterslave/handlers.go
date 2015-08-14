// handlers.go
package main

import (
	"encoding/json"
	"github.com/Pallinder/go-randomdata"
	"log"
	"net/http"
)

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

func ReadUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	var user string

	// Use slave to select
	if err := Conn.QueryRow("SELECT (user) FROM user.user WHERE id=?", id).Scan(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(user))
}

func GenUserHandler(w http.ResponseWriter, r *http.Request) {
	user := randomdata.SillyName()
	id := randomdata.Digits(10)

	//Use master to insert
	if _, err := Conn.Exec("INSERT INTO user.user (id,user) VALUES (?, ?)", id, user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(id))
}
