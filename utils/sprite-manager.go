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

	s.FrameSize.Left = cellWidth
	s.FrameSize.Right = cellHeight

	return s, nil
}

func (s *Sprite) SelectSubImage(col, row int) *ebiten.Image {
	x := (col - 1) * s.FrameSize.Left
	y := (row - 1) * s.FrameSize.Right

	return s.Image.SubImage(image.Rect(x, y, x+s.FrameSize.Left, y+s.FrameSize.Right)).(*ebiten.Image)
}
