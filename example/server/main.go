package main

import (
	"fmt"
	"log"

	"towerman1990.cn/zinx-learn/iface"
	"towerman1990.cn/zinx-learn/server"
	"towerman1990.cn/zinx-learn/utils"
)

type PingRouter struct {
	server.BaseRouter
}

func (pr *PingRouter) Handle(request iface.IRequest) {
	log.Println("call ping Handle function")
	log.Printf("receive message from client: message ID = [%d], message content = [%s]",
		request.GetMessageID(), string(request.GetMessageData()))
	request.GetConnection().SendMsg(100, []byte(fmt.Sprintf("recieved message [%d] successfully", request.GetMessageID())))
}

type HelloRouter struct {
	server.BaseRouter
}

func (pr *HelloRouter) Handle(request iface.IRequest) {
	log.Println("call hello Handle function")
	log.Printf("receive message from client: message ID = [%d], message content = [%s]",
		request.GetMessageID(), string(request.GetMessageData()))
	request.GetConnection().SendMsg(200, []byte(fmt.Sprintf("recieved message [%d] successfully", request.GetMessageID())))
}

func main() {
	s := server.New(utils.GlobalObject.Name)
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.Serve()
}
