package utils

import (
	"image"
	"math"
	"slices"

	"github.com/duckproblems/cogni/core/ecs/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SpriteMovementCollection struct {
	Up, Down, Left, Right []*ebiten.Image
}

type sprite struct {
	Image     *ebiten.Image
	FrameSize Pair[int]

	totalImages Pair[int]
}

func LoadSpriteFromPath(path string, cellWidth, cellHeight int) (*sprite, error) {
	image, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, err
	}

	s := &sprite{Image: image}

	s.FrameSize.Left = cellWidth
	s.FrameSize.Right = cellHeight

	return s, nil
}

func (s *sprite) SelectSubImage(col, row int) *ebiten.Image {
	x := (col - 1) * s.FrameSize.Left
	y := (row - 1) * s.FrameSize.Right

	return s.Image.SubImage(image.Rect(x, y, x+s.FrameSize.Left, y+s.FrameSize.Right)).(*ebiten.Image)
}

func (s *sprite) SelectFromRange(fromCol, fromRow, toCol, toRow int) []*ebiten.Image {
	total := (toCol - fromCol + 1) * (toRow - fromRow + 1)
	images := make([]*ebiten.Image, 0, total)

	for row := fromRow; row <= toRow; row++ {
		for col := fromCol; col <= toCol; col++ {
			images = append(images, s.SelectSubImage(col, row))
		}
	}

	return images
}

func UpdateMovementAnimationByDirection(sprite *components.Sprite, vX, vY float64, collection SpriteMovementCollection) {
	absX := math.Abs(vX)
	absY := math.Abs(vY)

	var targetFrames []*ebiten.Image

	if absX >= absY {
		if vX < 0 {
			targetFrames = collection.Left
		} else {
			targetFrames = collection.Right
		}
	} else {
		if vY < 0 {
			targetFrames = collection.Up
		} else {
			targetFrames = collection.Down
		}
	}

	if !slices.Equal(sprite.Frames, targetFrames) && targetFrames != nil {
		sprite.Frames = targetFrames
		sprite.CurrentFrame = 0
		sprite.Playing = true
	}
}
