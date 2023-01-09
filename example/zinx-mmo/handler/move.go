package handler

import (
	"log"

	"github.com/towerman1990/zinx-learn/iface"
	"github.com/towerman1990/zinx-learn/server"
	"google.golang.org/protobuf/proto"
	"towerman1990.cn/zinx-mmo/core"
	"towerman1990.cn/zinx-mmo/pb"
)

type MoveHandler struct {
	server.BaseRouter
}

func (wc *MoveHandler) Handle(request iface.IRequest) {
	protoMsg := &pb.Position{}
	err := proto.Unmarshal(request.GetMessageData(), protoMsg)
	if err != nil {
		log.Printf("move unmarshal request data failed, error: %s", err.Error())
	}

	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		log.Printf("get player pid failed, error: %s", err.Error())
	}

	log.Printf("player [%d] move: %f, %f, %f, %f", pid, protoMsg.X, protoMsg.Y, protoMsg.Z, protoMsg.V)

	player := core.WorldManagerObj.GetPlayerByPid(pid.(int32))
	player.UpdatePos(protoMsg.X, protoMsg.Y, protoMsg.Z, protoMsg.V)
}
