// config.go
package main

import (
	"github.com/coreos/go-etcd/etcd"
	"log"
	"strings"
)

var AppConfig *Config

type MysqlConfig struct {
	Host string
	Port string
	User string
	Pass string
}

type Config struct {
	Master MysqlConfig
	Slaves []MysqlConfig
}

func ReloadConfig(client *etcd.Client) {
	resp, err := client.Get(EtcdPath+"/master", false, true)
	if err != nil {
		panic(err)
	}
	setConfig(resp.Node.Nodes, &AppConfig.Master)

	resp, err = client.Get(EtcdPath+"/slaves", false, true)
	if err != nil {
		panic(err)
	}
	salvesConfig := []MysqlConfig{}
	for i := 0; i < len(resp.Node.Nodes); i++ {
		slaveConfig := MysqlConfig{}
		setConfig(resp.Node.Nodes[i].Nodes, &slaveConfig)
		salvesConfig = append(salvesConfig, slaveConfig)
	}
	AppConfig.Slaves = salvesConfig

	OpenSqlConnection()
}

func setConfig(nodes []*etcd.Node, config *MysqlConfig) {
	for i := 0; i < len(nodes); i++ {
		switch {
		case strings.HasSuffix(nodes[i].Key, "host"):
			config.Host = nodes[i].Value
		case strings.HasSuffix(nodes[i].Key, "port"):
			config.Port = nodes[i].Value
		case strings.HasSuffix(nodes[i].Key, "user"):
			config.User = nodes[i].Value
		case strings.HasSuffix(nodes[i].Key, "pass"):
			config.Pass = nodes[i].Value
		}
	}
}

func WatchConfig(client *etcd.Client) {
	watchChan := make(chan *etcd.Response)
	go client.Watch(EtcdPath, 0, true, watchChan, nil)
	go func() {
		for {
			select {
			case resp := <-watchChan:
				log.Printf("Config is changed, %s, %s", resp.Node.Key, resp.Node.Value)
				ReloadConfig(client)
			}
		}
	}()
}
