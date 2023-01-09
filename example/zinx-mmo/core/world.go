package core

import "sync"

type WorldManager struct {
	AoiMrg *AoiManager

	Players map[int32]*Player

	pLock sync.RWMutex
}

var WorldManagerObj *WorldManager

func init() {
	WorldManagerObj = &WorldManager{
		AoiMrg:  NewAoiManager(AOI_MIN_X, AOI_MAX_X, AOI_CNTS_X, AOI_MIN_Y, AOI_MAX_Y, AOI_CNTS_Y),
		Players: make(map[int32]*Player),
	}
}

func (wm *WorldManager) AddPlayer(player *Player) {
	wm.pLock.Lock()
	wm.Players[player.Pid] = player
	wm.pLock.Unlock()

	wm.AoiMrg.AddPlayerIdToGridByPos(int(player.Pid), player.X, player.Z)
}

func (wm *WorldManager) RemovePlayerByPid(pid int32) {
	player := wm.Players[pid]
	wm.AoiMrg.RemovePlayerIdByPosFromGrid(int(pid), player.X, player.Z)

	wm.pLock.Lock()
	delete(wm.Players, pid)
	wm.pLock.Unlock()
}

func (wm *WorldManager) GetPlayerByPid(pid int32) (player *Player) {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()
	return wm.Players[pid]
}

func (wm *WorldManager) GetAllPlayers() (players []*Player) {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()
	players = make([]*Player, 0)
	for _, player := range wm.Players {
		players = append(players, player)
	}

	return players
}
