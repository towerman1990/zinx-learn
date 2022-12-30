package iface

type IServer interface {
	Start()

	Serve()

	Stop()

	AddRouter(msgID uint32, router IRouter) error
}
