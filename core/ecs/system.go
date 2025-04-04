package ecs

import (
	"github.com/duckproblems/cogni/core/input"
	"github.com/hajimehoshi/ebiten/v2"
)

type System interface {
	Update(manger *ECSManager, inputManager *input.Manager, delta float64)
	Draw(manger *ECSManager, screen *ebiten.Image)
}
