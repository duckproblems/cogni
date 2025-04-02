package main

import (
	"github.com/duckproblems/cogni/core"
	"github.com/duckproblems/cogni/core/ecs"
	"github.com/duckproblems/cogni/core/ecs/components"
	"github.com/duckproblems/cogni/utils"
	"github.com/duckproblems/cogni/utils/graphics"
)

func main() {
	game := core.NewGame(*core.NewWindow(core.Window{Title: "Sentient", Geometry: utils.Pair[int]{Left: 800, Right: 600}}))

	playerSpriteSheet, _ := graphics.LoadSpriteSheet("./assets/sprites/player.png", 16, 16)
	playerAnimSet := graphics.NewAnimationSet()
	playerAnimSet.AddAnimation(graphics.AnimStateWalkUp, playerSpriteSheet.GetFrames(2, 1, 2, 4))
	playerAnimSet.AddAnimation(graphics.AnimStateWalkDown, playerSpriteSheet.GetFrames(1, 1, 1, 4))
	playerAnimSet.AddAnimation(graphics.AnimStateWalkLeft, playerSpriteSheet.GetFrames(3, 1, 3, 4))
	playerAnimSet.AddAnimation(graphics.AnimStateWalkRight, playerSpriteSheet.GetFrames(4, 1, 4, 4))

	player := ecs.NewEntity("Player")
	player.AddComponent(&components.Sprite{
		Frames:     playerSpriteSheet.GetFrames(1, 1, 1, 1),
		FrameSpeed: 5.,
		Loop:       true,
		Playing:    true,
	})
	player.AddComponent(&components.Transform{X: 30, Y: 30, ScaleX: 5, ScaleY: 5, Rotation: 0})
	player.AddComponent(&components.Input{})

	var sprite *components.Sprite
	player.GetComponent(&sprite)

	player.AddComponent(&components.Movement{
		MaxSpeed:     100.,
		Acceleration: 1000.,
		Friction:     0,
		OnStartedMoving: func(vX, vY float64) {
			graphics.UpdateMovementAnimation(sprite, playerAnimSet, vX, vY)
		},
		OnStoppedMoving: func(vX, vY float64) {
			graphics.UpdateMovementAnimation(sprite, playerAnimSet, vX, vY)
		},
		OnCruising: func(vX, vY float64) {
			graphics.UpdateMovementAnimation(sprite, playerAnimSet, vX, vY)
		},
	})

	game.ECS.AddEntity(player)
	game.Run()
}
