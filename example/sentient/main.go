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
	player.AddComponent(&components.Transform{X: 30, Y: 30, ScaleX: 1, ScaleY: 1, Rotation: 0})

	var sprite *components.Sprite
	player.GetComponent(&sprite)

	player.AddComponent(&components.InputControlled{})
	player.AddComponent(&components.MovementIntent{})
	player.AddComponent(&components.Movement{
		MaxSpeed:     50,
		Acceleration: 800,
		Friction:     1,
		OnStartedMoving: func(entity ecs.Entity, vX float64, vY float64) {
			graphics.UpdateMovementAnimation(sprite, playerAnimSet, vX, vY)
		},
		OnCruising: func(entity ecs.Entity, vX float64, vY float64) {
			graphics.UpdateMovementAnimation(sprite, playerAnimSet, vX, vY)
		},
		OnStoppedMoving: func(entity ecs.Entity, vX float64, vY float64) {
			graphics.UpdateMovementAnimation(sprite, playerAnimSet, vX, vY)
		},
	})

	npc := ecs.NewEntity("NPC")
	npc.AddComponent(&components.Sprite{
		Frames:     playerSpriteSheet.GetFrames(1, 1, 1, 1),
		FrameSpeed: 5.,
		Loop:       true,
		Playing:    true,
	})
	npc.AddComponent(&components.Transform{X: 100, Y: 30, ScaleX: 1, ScaleY: 1, Rotation: 0})

	npc.AddComponent(&components.MovementIntent{})
	npc.AddComponent(&components.Movement{
		MaxSpeed:     50,
		Acceleration: 800,
		Friction:     1,
	})

	game.ECS.AddEntity(player)
	game.ECS.AddEntity(npc)
	game.Run()
}
