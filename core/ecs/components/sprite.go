package components

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	Frames       []*ebiten.Image
	CurrentFrame int
	FrameSpeed   float64
	FrameTimer   float64
	Loop         bool
	Playing      bool
}
