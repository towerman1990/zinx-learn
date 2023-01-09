package core

import (
	"log"
	"math/rand"
	"sync"

	"github.com/towerman1990/zinx-learn/iface"
	"google.golang.org/protobuf/proto"
	pb "towerman1990.cn/zinx-mmo/pb"
)

var PidGen int32 = 1

var PidLock sync.Mutex

type Player struct {
	Pid  int32
	Conn iface.IConnection

	X float32
	Y float32
	Z float32
	V float32
}

func (p *Player) SendMsg(msgId uint32, data proto.Message) {
	msg, err := proto.Marshal(data)
	if err != nil {
		log.Printf("marshal msg failed, error: %s", err.Error())
		return
	}

	if p.Conn == nil {
		log.Println("connection in player is nil")
		return
	}

	if err := p.Conn.SendMsg(msgId, msg); err != nil {
		log.Printf("player send message failed, error: %s", err.Error())
		return
	}
}

func (p *Player) SyncPid() {
	data := &pb.SyncPid{
		Pid: p.Pid,
	}
	p.SendMsg(1, data)
}

func (p *Player) BroadCastStartPositon() {
	data := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2,
		Data: &pb.BroadCast_Pos{
			Pos: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	p.SendMsg(200, data)
}

func (p *Player) Talk(content string) {
	protoMsg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  1,
		Data: &pb.BroadCast_Content{
			Content: content,
		},
	}

	players := WorldManagerObj.GetAllPlayers()

	for _, player := range players {
		player.SendMsg(200, protoMsg)
	}
}

func (p *Player) SyncPosToSurroundPlayer() {
	pids := WorldManagerObj.AoiMrg.GetSurroundPlayerIds(p.X, p.Z)
	log.Println("surround pids:", pids)
	players := make([]*Player, 0, len(pids))
	for _, pid := range pids {
		players = append(players, WorldManagerObj.GetPlayerByPid(int32(pid)))
	}

	protoMsg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2,
		Data: &pb.BroadCast_Pos{
			Pos: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	log.Println("surround players:", players)
	for _, player := range players {
		player.SendMsg(200, protoMsg)
	}

	playersProtoMsg := make([]*pb.Player, 0, len(players))
	for _, player := range players {
		p := &pb.Player{
			Pid: player.Pid,
			Pos: &pb.Position{
				X: player.X,
				Y: player.Y,
				Z: player.Z,
				V: player.V,
			},
		}
		playersProtoMsg = append(playersProtoMsg, p)
	}

	syncPlayersProtoMsg := &pb.SyncPlayers{
		P: playersProtoMsg,
	}

	p.SendMsg(202, syncPlayersProtoMsg)
}

func (p *Player) UpdatePos(x, y, z, v float32) {
	p.X = x
	p.Y = y
	p.Z = z
	p.V = v

	protoMsg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  4,
		Data: &pb.BroadCast_Pos{
			Pos: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	players := p.GetSurroundPlayers()
	for _, player := range players {
		player.SendMsg(200, protoMsg)
	}
}

func (p *Player) GetSurroundPlayers() []*Player {
	pids := WorldManagerObj.AoiMrg.GetSurroundPlayerIds(p.X, p.Z)

	players := make([]*Player, 0, len(pids))
	for _, pid := range pids {
		players = append(players, WorldManagerObj.GetPlayerByPid(int32(pid)))
	}

	return players
}

func (p *Player) Offline() {
	protoMsg := &pb.SyncPid{
		Pid: p.Pid,
	}

	players := p.GetSurroundPlayers()
	for _, player := range players {
		player.SendMsg(201, protoMsg)
	}

	WorldManagerObj.AoiMrg.RemovePlayerIdByPosFromGrid(int(p.Pid), p.X, p.Z)
	WorldManagerObj.RemovePlayerByPid(p.Pid)
}

func NewPlayer(conn iface.IConnection) (player *Player) {
	PidLock.Lock()
	pid := PidGen
	PidGen++
	PidLock.Unlock()
	player = &Player{
		Pid:  pid,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)),
		Y:    0,
		Z:    float32(140 + rand.Intn(10)),
		V:    0,
	}

	return player
}
