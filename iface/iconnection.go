package iface

import "net"

type IConnecton interface {
	Open()

	Close() error

	GetTCPConnection() *net.TCPConn

	GetConnID() uint32

	RemoteAddr() net.Addr

	SendMsg(MsgID uint32, data []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
