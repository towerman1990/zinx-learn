package main

import (
	"fmt"
	"log"

	"github.com/towerman1990/zinx-learn/iface"
	"github.com/towerman1990/zinx-learn/server"
	"github.com/towerman1990/zinx-learn/utils"
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

func CallOnConnectionOpen(conn iface.IConnection) {
	log.Println("CallOnConnectionOpen is called")
	conn.SendMsg(200, []byte("CallOnConnectionOpen is called"))

	conn.SetProperty("name", "andy")
	conn.SetProperty("age", "22")
}

func CallOnConnectionClose(conn iface.IConnection) {
	log.Println("CallOnConnectionClose is called")
	if property, err := conn.GetProperty("name"); err == nil {
		log.Printf("name = %s", property)
	} else {
		log.Printf("get property name failed, error: %s", err.Error())
	}

	if property, err := conn.GetProperty("age"); err == nil {
		log.Printf("age = %s", property)
	} else {
		log.Printf("get property age failed, error: %s", err.Error())
	}

	if property, err := conn.GetProperty("mobile"); err == nil {
		log.Printf("mobile = %s", property)
	} else {
		log.Printf("get property mobile failed, error: %s", err.Error())
	}
}

func main() {
	s := server.New(utils.GlobalObject.Name)
	s.SetOnConnOpen(CallOnConnectionOpen)
	s.SetOnConnClose(CallOnConnectionClose)

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.Serve()
}
