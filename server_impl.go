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

			go func() {
				for {
					buf := make([]byte, 512)
					if _, err := conn.Read(buf); err != nil {
						log.Printf("conn read data failed, error: %s", err.Error())
						continue
					}

					log.Printf("server receive message: %s", buf)

					if _, err = conn.Write(buf); err != nil {
						log.Printf("conn write data failed, error: %s", err.Error())
						continue
					}

				}
			}()
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
