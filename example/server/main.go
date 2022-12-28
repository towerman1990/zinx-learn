package main

import (
	"log"

	iface "towerman1990.cn/zinx-learn/iface"
	"towerman1990.cn/zinx-learn/server"
)

type PingRouter struct {
	server.BaseRouter
}

func (pr *PingRouter) PreHandle(request iface.IRequest) {
	log.Println("call PreHandle function")
	if _, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n")); err != nil {
		log.Printf("PreHandle function write data failed, error: %s", err.Error())
	}
}

func (pr *PingRouter) Handle(request iface.IRequest) {
	log.Println("call Handle function")
	if _, err := request.GetConnection().GetTCPConnection().Write([]byte("is pinging...\n")); err != nil {
		log.Printf("Handle function write data failed, error: %s", err.Error())
	}
}

func (pr *PingRouter) PostHandle(request iface.IRequest) {
	log.Println("call PostHandle function")
	if _, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...")); err != nil {
		log.Printf("PostHandle function write data failed, error: %s", err.Error())
	}
}

func main() {
	s := server.New("zinx learn")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
