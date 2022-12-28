package iface

type IServer interface {
	Start()

	Serve()

	Stop()

	AddRouter(router IRouter)
}
