package core

import (
	"fmt"
	"math"
	"time"

	"github.com/duckproblems/cogni/core/ecs"
	"github.com/duckproblems/cogni/core/ecs/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Window *Window
	ECS    *ecs.ECSManager

	lastUpdate time.Time
}

func NewGame(window Window) *Game {
	ecsManager := ecs.New()
	ecsManager.AddSystem(systems.Animate{})

	return &Game{
		Window:     &window,
		ECS:        ecsManager,
		lastUpdate: time.Now(),
	}
}

func (g *Game) Run() error {
	return ebiten.RunGame(g)
}

func (g *Game) Update() error {
	delta := time.Since(g.lastUpdate).Seconds()
	g.lastUpdate = time.Now()

	g.ECS.Update(delta)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ECS.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %d", int(math.Floor(ebiten.ActualFPS()))))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}
