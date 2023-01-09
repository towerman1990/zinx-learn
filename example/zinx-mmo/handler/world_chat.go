package handler

import (
	"log"

	"github.com/towerman1990/zinx-learn/iface"
	"github.com/towerman1990/zinx-learn/server"
	"google.golang.org/protobuf/proto"
	"towerman1990.cn/zinx-mmo/core"
	"towerman1990.cn/zinx-mmo/pb"
)

type WorldChatHandler struct {
	server.BaseRouter
}

func (wc *WorldChatHandler) Handle(request iface.IRequest) {
	protoMsg := &pb.Talk{}
	err := proto.Unmarshal(request.GetMessageData(), protoMsg)
	if err != nil {
		log.Printf("talk unmarshal request data failed, error: %s", err.Error())
	}

	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		log.Printf("get player pid failed, error: %s", err.Error())
	}

	player := core.WorldManagerObj.GetPlayerByPid(pid.(int32))
	player.Talk(protoMsg.Content)
}
