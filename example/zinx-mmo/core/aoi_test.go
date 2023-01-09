package core

import (
	"testing"
)

func TestNewAoiManger(t *testing.T) {
	aoiManager := NewAoiManager(0, 500, 5, 0, 400, 4)
	t.Log(aoiManager.String())
}

func TestGetSurroundGrids(t *testing.T) {
	aoiManager := NewAoiManager(0, 500, 5, 0, 400, 4)
	grids := aoiManager.GetSurroundGrids(0)
	for _, grid := range grids {
		t.Log(grid.String())
	}
}
