package cogni

import (
	"fmt"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type gameManager struct {
	Window *windowManager
	Input  *inputManager

	lastUpdate time.Time
}

func New(title string, width, height int) *gameManager {
	return &gameManager{
		Window:     newWindowManager(title, width, height),
		Input:      newInputManager(),
		lastUpdate: time.Now(),
	}
}

func (g *gameManager) Run() error {
	return ebiten.RunGame(g)
}

func (g *gameManager) Update() error {
	delta := time.Since(g.lastUpdate).Seconds()
	g.lastUpdate = time.Now()

	for _, key := range inpututil.AppendJustPressedKeys(nil) {
		g.Input.execPressEvent(key, delta)
	}

	for _, key := range inpututil.AppendPressedKeys(nil) {
		g.Input.execHoldEvent(key, delta)
	}

	for _, key := range inpututil.AppendJustReleasedKeys(nil) {
		g.Input.execReleaseEvent(key, delta)
	}

	return nil
}

func (g *gameManager) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %d", int(math.Floor(ebiten.ActualFPS()))))
}

func (g *gameManager) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}
