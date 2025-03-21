package cogni

import "github.com/hajimehoshi/ebiten/v2"

type windowManager struct {
	title         string
	width, height int
}

func newWindowManager(title string, width, height int) *windowManager {
	return &windowManager{title, width, height}
}

func (w *windowManager) SetTitle(newTitle string) {
	w.title = newTitle

	ebiten.SetWindowTitle(w.title)
}

func (w *windowManager) Resize(width, height int) {
	w.width = width
	w.height = height

	ebiten.SetWindowSize(w.width, w.height)
}
