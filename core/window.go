package core

import (
	"github.com/duckproblems/cogni/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Window struct {
	Title    string
	Geometry utils.Pair[int]
}

func NewWindow(config Window) *Window {
	win := &Window{}

	win.WithTitle(config.Title)
	win.WithGeometry(config.Geometry.Left, config.Geometry.Right)

	return win
}

func (w *Window) WithTitle(title string) *Window {
	w.Title = title

	ebiten.SetWindowTitle(title)

	return w
}

func (w *Window) WithGeometry(width, height int) *Window {
	w.Geometry.Left = width
	w.Geometry.Right = height

	ebiten.SetWindowSize(width, height)

	return w
}
