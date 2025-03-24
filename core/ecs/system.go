package ecs

import "github.com/hajimehoshi/ebiten/v2"

type System interface {
	Update(manger *ECSManager, delta float64)
	Draw(manger *ECSManager, screen *ebiten.Image)
}
