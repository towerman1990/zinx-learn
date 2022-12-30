package server

import (
	"fmt"
	"log"
	"net"

	"towerman1990.cn/zinx-learn/iface"
	"towerman1990.cn/zinx-learn/utils"
)

type Server struct {
	Name string

	IPVersion string

	IP string

	Port int

	MessageHandler iface.IMessageHandler
}

var GlobalConnID uint32 = 0

func MirrorRespond(conn *net.TCPConn, data []byte, dataLen int) (err error) {
	if _, err = conn.Write(data[:dataLen]); err != nil {
		log.Printf("[MirrorRespond] conn write data failed, error: %s", err.Error())
	}
	return
}

func (s *Server) Start() {
	log.Printf("server [%s] start", s.Name)
	go func() {
		s.MessageHandler.StartWorkPool()

		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			log.Fatalf("server [%s] resolve TCP addr failed, error: %s", s.Name, err.Error())
		}

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			log.Fatalf("server [%s] listened TCP failed, error: %s", s.Name, err.Error())
		}

		log.Printf("server [%s] is listenning at IP: [%s], Port: [%d]", s.Name, s.IP, s.Port)

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Printf("listener accept TCP failed, error: %s", err.Error())
				continue
			}

			connID := GlobalConnID
			dealConn := NewConnection(conn, connID, s.MessageHandler)
			GlobalConnID++

			go dealConn.Open()
		}
	}()
}

func (s *Server) Serve() {
	s.Start()
	select {}
}

func (s *Server) Stop() {

}

func (s *Server) AddRouter(msgID uint32, router iface.IRouter) error {
	return s.MessageHandler.AddRouter(msgID, router)
}

func New(name string) iface.IServer {
	s := &Server{
		Name:           utils.GlobalObject.Name,
		IPVersion:      "tcp4",
		IP:             utils.GlobalObject.Host,
		Port:           utils.GlobalObject.Port,
		MessageHandler: NewMessageHandler(),
	}

	return s
}
