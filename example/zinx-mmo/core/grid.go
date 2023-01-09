package core

import (
	"fmt"
	"log"
	"sync"
)

type Grid struct {
	ID int

	MinX int

	MaxX int

	MinY int

	MaxY int

	playerIdMap map[int]bool

	Lock sync.RWMutex
}

func NewGrid(gridId, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		ID:          gridId,
		MinX:        minX,
		MaxX:        maxX,
		MinY:        minY,
		MaxY:        maxY,
		playerIdMap: make(map[int]bool),
	}
}

func (g *Grid) AddPlayer(playerId int) {
	log.Printf("g: %v", g)
	g.Lock.Lock()
	defer g.Lock.Unlock()
	g.playerIdMap[playerId] = true
}

func (g *Grid) RemovePlayer(playerId int) {
	g.Lock.Lock()
	defer g.Lock.Unlock()

	delete(g.playerIdMap, playerId)
}

func (g *Grid) GetPlayerIds() (playerIds []int) {
	g.Lock.RLock()
	defer g.Lock.RUnlock()

	for pid := range g.playerIdMap {
		playerIds = append(playerIds, pid)
	}

	return playerIds
}

func (g *Grid) String() string {
	return fmt.Sprintf("Grid ID: %d, minX: %d, maxX: %d, minY: %d, maxY: %d, playerIds: %v\n",
		g.ID, g.MinX, g.MaxX, g.MinY, g.MaxY, g.playerIdMap)
}
