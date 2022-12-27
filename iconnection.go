package zinxlearn

import "net"

type IConnecton interface {
	Open()

	Close()

	GetTCPConnection() *net.TCPConn

	GetConnID() uint32

	RemoteAddr() net.Addr

	Send(data []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
