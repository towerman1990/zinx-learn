package iface

type IServer interface {
	Start()

	Serve()

	Stop()

	AddRouter(msgID uint32, router IRouter) error

	GetConnManager() IConnManager

	SetOnConnOpen(func(conn IConnection))

	SetOnConnClose(func(conn IConnection))

	CallOnConnOpen(conn IConnection)

	CallOnConnClose(conn IConnection)
}
