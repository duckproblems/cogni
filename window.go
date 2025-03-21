package cogni

import "github.com/hajimehoshi/ebiten/v2"

type windowManager struct {
	title         string
	width, height int
}

func newWindowManager(title string, width, height int) *windowManager {
	wm := windowManager{}

	wm.SetTitle(title)
	wm.Resize(width, height)

	return &wm
}

func (w *windowManager) SetTitle(newTitle string) {
	w.title = newTitle

	ebiten.SetWindowTitle(w.title)
}

func (w *windowManager) Resize(newWidth, newHeight int) {
	w.width = newWidth
	w.height = newHeight

	ebiten.SetWindowSize(w.width, w.height)
}
