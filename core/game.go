package core

import (
	"fmt"
	"math"
	"time"

	"github.com/duckproblems/cogni/core/clog"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Logger *clog.Logger
	Window *Window

	lastUpdate time.Time
}

func NewGame(window Window) *Game {
	return &Game{
		Window:     &window,
		lastUpdate: time.Now(),
	}
}

func (g *Game) Run() error {
	return ebiten.RunGame(g)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %d", int(math.Floor(ebiten.ActualFPS()))))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}
