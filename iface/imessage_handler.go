package iface

type IMessageHandler interface {
	ExecHandler(request IRequest)

	AddRouter(msgID uint32, router IRouter) error

	StartWorkPool()

	SendMsgToTaskQueue(IRequest)
}
