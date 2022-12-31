package iface

import "net"

type IConnection interface {
	Open()

	Close() error

	GetTCPConnection() *net.TCPConn

	GetConnID() uint32

	RemoteAddr() net.Addr

	SendMsg(MsgID uint32, data []byte) error

	SetProperty(string, interface{})

	GetProperty(key string) (interface{}, error)

	RemoveProperty(key string)
}

type HandleFunc func(*net.TCPConn, []byte, int) error
