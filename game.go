package cogni

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type gameManager struct {
	Window *windowManager
	Input  *inputManager
}

func New(title string, width, height int) *gameManager {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)

	return &gameManager{
		Window: newWindowManager(title, width, height),
		Input:  newInputManager(),
	}
}

func (g *gameManager) Run() error {
	return ebiten.RunGame(g)
}

func (g *gameManager) Update() error {
	for _, key := range inpututil.AppendJustPressedKeys(nil) {
		g.Input.execPressEvent(key)
	}

	for _, key := range inpututil.AppendPressedKeys(nil) {
		g.Input.execHoldEvent(key)
	}

	for _, key := range inpututil.AppendJustReleasedKeys(nil) {
		g.Input.execReleaseEvent(key)
	}

	return nil
}

func (g *gameManager) Draw(screen *ebiten.Image) {
}

func (g *gameManager) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Window.width, g.Window.height
}
