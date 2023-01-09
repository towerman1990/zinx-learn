package server

import (
	"fmt"
	"log"

	"github.com/towerman1990/zinx-learn/iface"
	"github.com/towerman1990/zinx-learn/utils"
)

type MessageHandler struct {
	Handlers       map[uint32]iface.IRouter
	TaskQueue      []chan iface.IRequest
	WorkerPoolSize uint32
}

func (mh *MessageHandler) ExecHandler(request iface.IRequest) {
	msgID := request.GetMessageID()
	handler, ok := mh.Handlers[msgID]
	if !ok {
		fmt.Printf("msgID [%d] hasn't been added\n", msgID)
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

func (mh *MessageHandler) StartWorkPool() {
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		mh.TaskQueue[i] = make(chan iface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		go mh.StartOneWork(i)
	}
}

func (mh *MessageHandler) StartOneWork(i int) {
	log.Printf("worker [%d] started", i)

	for request := range mh.TaskQueue[i] {
		mh.ExecHandler(request)
	}
}

func (mh *MessageHandler) SendMsgToTaskQueue(request iface.IRequest) {
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	mh.TaskQueue[workerID] <- request
}

func NewMessageHandler() iface.IMessageHandler {
	return &MessageHandler{
		Handlers:       make(map[uint32]iface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan iface.IRequest, utils.GlobalObject.MaxWorkerTaskLen),
	}
}
