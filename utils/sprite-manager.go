package utils

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Image     *ebiten.Image
	FrameSize Pair[int]

	totalImages Pair[int]
}

func LoadFromPath(path string, cellWidth, cellHeight int) (*Sprite, error) {
	sprite, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, err
	}

	s := &Sprite{Image: sprite}

	s.FrameSize.A = cellWidth
	s.FrameSize.B = cellHeight

	return s, nil
}

func (s *Sprite) SelectSubImage(col, row int) *ebiten.Image {
	x := (col - 1) * s.FrameSize.A
	y := (row - 1) * s.FrameSize.B

	return s.Image.SubImage(image.Rect(x, y, x+s.FrameSize.A, y+s.FrameSize.B)).(*ebiten.Image)
}
