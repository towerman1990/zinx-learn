package main

import (
	"fmt"
	"log"

	iface "towerman1990.cn/zinx-learn/iface"
	"towerman1990.cn/zinx-learn/server"
)

type PingRouter struct {
	server.BaseRouter
}

func (pr *PingRouter) Handle(request iface.IRequest) {
	log.Println("call Handle function")
	log.Printf("receive message from client: message ID = [%d], message content = [%s]",
		request.GetMessageID(), string(request.GetMessageData()))
	request.GetConnection().SendMsg(1001, []byte(fmt.Sprintf("recieved message [%d] successfully", request.GetMessageID())))
}

func main() {
	s := server.New("zinx learn")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
