package main

import (
	"log"

	"github.com/towerman1990/zinx-learn/iface"
	"github.com/towerman1990/zinx-learn/server"
	"towerman1990.cn/zinx-mmo/core"
	"towerman1990.cn/zinx-mmo/handler"
)

func main() {
	s := server.New("Zinx MMO Game")
	s.SetOnConnOpen(OnConnectionAdd)
	s.SetOnConnClose(OnConnectionLost)

	s.AddRouter(2, &handler.WorldChatHandler{})
	s.AddRouter(3, &handler.MoveHandler{})

	s.Serve()
}

func OnConnectionAdd(conn iface.IConnection) {
	player := core.NewPlayer(conn)

	player.SyncPid()

	player.BroadCastStartPositon()

	core.WorldManagerObj.AddPlayer(player)

	conn.SetProperty("pid", player.Pid)

	player.SyncPosToSurroundPlayer()

	log.Printf("player [%d] join world", player.Pid)
}

func OnConnectionLost(conn iface.IConnection) {
	pid, err := conn.GetProperty("pid")
	if err != nil {
		log.Printf("get player pid failed, error: %s", err.Error())
	}

	player := core.WorldManagerObj.GetPlayerByPid(pid.(int32))
	player.Offline()

	log.Printf("player [%d] leave world", player.Pid)
}
