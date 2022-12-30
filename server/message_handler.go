package server

import (
	"fmt"
	"log"

	"towerman1990.cn/zinx-learn/iface"
)

type MessageHandler struct {
	Handlers map[uint32]iface.IRouter
}

func (mh *MessageHandler) ExecHandler(request iface.IRequest) {
	msgID := request.GetMessageID()
	handler, ok := mh.Handlers[msgID]
	if !ok {
		fmt.Printf("msgID [%d] hasn't been added", msgID)
		return
	}

	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (mh *MessageHandler) AddRouter(msgID uint32, router iface.IRouter) (err error) {
	if _, ok := mh.Handlers[msgID]; ok {
		return fmt.Errorf("msgID [%d] has been added", msgID)
	}

	mh.Handlers[msgID] = router
	log.Printf("added router successfully, msgID = [%d]", msgID)

	return
}

func NewMessageHandler() iface.IMessageHandler {
	return &MessageHandler{
		Handlers: make(map[uint32]iface.IRouter),
	}
}
