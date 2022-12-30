package utils

import (
	"encoding/json"
	"log"
	"os"

	"towerman1990.cn/zinx-learn/iface"
)

type GlobalObj struct {
	TcpServer iface.IServer

	Host string
	Port int
	Name string

	Version          string
	MaxConn          int
	MaxPackageSize   uint32
	WorkerPoolSize   uint32
	MaxWorkerTaskLen uint32
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() (err error) {
	data, err := os.ReadFile("conf/zinx.json")
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &GlobalObject)
}

func init() {
	GlobalObject = &GlobalObj{
		Name:             "Zinx Learn",
		Version:          "v0.4",
		Port:             2022,
		Host:             "0.0.0.0",
		MaxConn:          1024,
		MaxPackageSize:   4096,
		WorkerPoolSize:   12,
		MaxWorkerTaskLen: 1024,
	}

	if err := GlobalObject.Reload(); err != nil {
		log.Printf("reload config file failed, error: %s", err.Error())
	}
}
