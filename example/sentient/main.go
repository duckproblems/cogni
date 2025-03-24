package main

import (
	"github.com/duckproblems/cogni/core"
	"github.com/duckproblems/cogni/core/ecs"
	"github.com/duckproblems/cogni/core/ecs/components"
	"github.com/duckproblems/cogni/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := core.NewGame(*core.NewWindow(core.Window{Title: "Sentient", Geometry: utils.Pair[int]{Left: 800, Right: 600}}))

	playerImage, _ := utils.LoadFromPath("./assets/sprites/player.png", 16, 16)

	player := ecs.NewEntity("Player")
	player.AddComponent(&components.Sprite{
		Frames: []*ebiten.Image{
			playerImage.SelectSubImage(1, 1),
			playerImage.SelectSubImage(1, 2),
			playerImage.SelectSubImage(1, 3),
			playerImage.SelectSubImage(1, 4),
		},
		FrameSpeed: 5.,
		Loop:       true,
		Playing:    true,
	})
	player.AddComponent(&components.Transform{X: 30, Y: 30, ScaleX: 1, ScaleY: 1, Rotation: 0})

	game.ECS.AddEntity(player)
	game.Run()
}
