package zinxlearn

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	Name string

	IPVersion string

	IP string

	Port int
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
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			log.Fatalf("server [%s] resolve TCP addr failed, error: %s", s.Name, err.Error())
		}

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			log.Fatalf("server [%s] listen TCP failed, error: %s", s.Name, err.Error())
		}

		log.Printf("server [%s] started successfully and it's listenning at IP: [%s], Port: [%d]", s.Name, s.IP, s.Port)

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Printf("listener accept TCP failed, error: %s", err.Error())
				continue
			}

			connID := GlobalConnID
			dealConn := NewConnection(conn, connID, MirrorRespond)
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

func New(name string) IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8888,
	}

	return s
}
