package graphics

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SpriteSheet struct {
	image *ebiten.Image

	frameWidth, frameHeight int
}

func LoadSpriteSheet(path string, frameWidth, frameHeight int) (*SpriteSheet, error) {
	if frameWidth <= 0 || frameHeight <= 0 {
		return nil, fmt.Errorf("invalid frame dimensions: width=%d, height=%d must be positive", frameWidth, frameHeight)
	}

	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, err
	}

	sheet := &SpriteSheet{
		image:       img,
		frameWidth:  frameWidth,
		frameHeight: frameHeight,
	}

	return sheet, nil
}

func (s *SpriteSheet) GetFrame(col, row int) *ebiten.Image {
	if col < 1 || row < 1 {
		return nil
	}

	col0 := col - 1
	row0 := row - 1

	x := col0 * s.frameWidth
	y := row0 * s.frameHeight

	imgBounds := s.image.Bounds()
	if x >= imgBounds.Dx() || y >= imgBounds.Dy() {
		return nil
	}

	rect := image.Rect(x, y, x+s.frameWidth, y+s.frameHeight)
	subRect := rect.Intersect(imgBounds)
	if subRect.Empty() || subRect.Dx() != s.frameWidth || subRect.Dy() != s.frameHeight {
		return nil
	}

	subImg, ok := s.image.SubImage(rect).(*ebiten.Image)
	if !ok {
		return nil
	}

	return subImg
}

func (s *SpriteSheet) GetFrames(fromCol, fromRow, toCol, toRow int) []*ebiten.Image {
	if fromCol > toCol || fromRow > toRow || fromCol < 1 || fromRow < 1 {
		return []*ebiten.Image{}
	}

	numCols := toCol - fromCol + 1
	numRows := toRow - fromRow + 1
	estimatedTotal := numCols * numRows

	images := make([]*ebiten.Image, 0, estimatedTotal)

	for r := fromRow; r <= toRow; r++ {
		for c := fromCol; c <= toCol; c++ {
			frame := s.GetFrame(c, r)
			if frame != nil {
				images = append(images, frame)
			}
		}
	}

	return images
}
