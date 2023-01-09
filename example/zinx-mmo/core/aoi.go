package core

import (
	"fmt"
	"log"
)

const (
	AOI_MIN_X  int = 85
	AOI_MAX_X  int = 410
	AOI_CNTS_X int = 10
	AOI_MIN_Y  int = 75
	AOI_MAX_Y  int = 400
	AOI_CNTS_Y int = 20
)

type AoiManager struct {
	MinX int

	MaxX int

	GridCountX int

	MinY int

	MaxY int

	GridCountY int

	GridMap map[int]*Grid
}

func (am *AoiManager) gridWidth() int {
	return (am.MaxX - am.MinX) / am.GridCountX
}

func (am *AoiManager) gridLength() int {
	return (am.MaxY - am.MinY) / am.GridCountY
}

func (am *AoiManager) GetSurroundGrids(gridID int) (grids []*Grid) {
	if _, ok := am.GridMap[gridID]; !ok {
		return
	}

	grids = append(grids, am.GridMap[gridID])

	idx := gridID % am.GridCountX

	if idx > 0 {
		grids = append(grids, am.GridMap[gridID-1])
	}

	if idx < am.GridCountX-1 {
		grids = append(grids, am.GridMap[gridID+1])
	}

	gidsX := make([]int, 0, len(grids))
	for _, v := range grids {
		gidsX = append(gidsX, v.ID)
	}

	for _, v := range gidsX {
		idy := v / am.GridCountY
		if idy > 0 {
			grids = append(grids, am.GridMap[v-am.GridCountX])
		}

		if idy < am.GridCountY-1 {
			grids = append(grids, am.GridMap[v+am.GridCountX])
		}
	}

	return
}

func (am *AoiManager) GetGridIdByPos(x, y float32) (gridId int) {
	idx := (int(x) - am.MinX) / am.gridWidth()
	idy := (int(y) - am.MinY) / am.gridLength()

	return idy*am.GridCountX + idx
}

func (am *AoiManager) GetSurroundPlayerIds(x, y float32) (playerIds []int) {
	gridId := am.GetGridIdByPos(x, y)
	grids := am.GetSurroundGrids(gridId)
	for _, grid := range grids {
		playerIds = append(playerIds, grid.GetPlayerIds()...)
	}

	return playerIds
}

func (am *AoiManager) AddPlayerIdToGrid(playerId, gridId int) {
	am.GridMap[gridId].AddPlayer(playerId)
}

func (am *AoiManager) RemovePlayerIdFromGrid(playerId, gridId int) {
	am.GridMap[gridId].RemovePlayer(playerId)
}

func (am *AoiManager) GetPlayerIdsFromGrid(gridId int) (playerIds []int) {
	playerIds = am.GridMap[gridId].GetPlayerIds()
	return
}

func (am *AoiManager) AddPlayerIdToGridByPos(playerId int, x, y float32) {
	gridId := am.GetGridIdByPos(x, y)
	grid := am.GridMap[gridId]
	grid.AddPlayer(playerId)
}

func (am *AoiManager) RemovePlayerIdByPosFromGrid(playerId int, x, y float32) {
	gridId := am.GetGridIdByPos(x, y)
	grid := am.GridMap[gridId]
	grid.RemovePlayer(playerId)
}

func (am *AoiManager) String() string {
	s := fmt.Sprintf("AoiManager:\n MinX: %d, MaxX: %d, GridCountX: %d, MinY: %d, MaxY: %d, GridCountY: %d\n",
		am.MinX, am.MaxX, am.GridCountX, am.MinY, am.MaxY, am.GridCountY,
	)

	for _, grid := range am.GridMap {
		s += grid.String()
	}

	return s
}

func NewAoiManager(minX, maxX, gridCountX, minY, maxY, gridCountY int) (aoiManager *AoiManager) {
	aoiManager = &AoiManager{
		MinX:       minX,
		MaxX:       maxX,
		GridCountX: gridCountX,
		MinY:       minY,
		MaxY:       maxY,
		GridCountY: gridCountY,
		GridMap:    make(map[int]*Grid),
	}

	for y := 0; y < gridCountY; y++ {
		for x := 0; x < gridCountX; x++ {
			gid := y*gridCountX + x

			aoiManager.GridMap[gid] = NewGrid(gid,
				aoiManager.MinX+x*aoiManager.gridWidth(),
				aoiManager.MinX+(x+1)*aoiManager.gridWidth(),
				aoiManager.MinY+y*aoiManager.gridLength(),
				aoiManager.MinY+(y+1)*aoiManager.gridLength(),
			)
		}
	}
	log.Printf("aoiManager: %v", aoiManager)

	return aoiManager
}
