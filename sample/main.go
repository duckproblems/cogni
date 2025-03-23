package main

import (
	"log"

	"github.com/duckproblems/cogni"
	"github.com/duckproblems/cogni/ecs/components"
	"github.com/duckproblems/cogni/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := cogni.New("Cogni Sample Window", 1024, 768)
	player := game.ECS.CreateEntity()

	playerImage, _ := utils.LoadFromPath("./assets/sprites/player.png", 16, 16)

	game.ECS.AddComponent(player, &components.Transform{X: 100, Y: 100, Rotation: 0})
	game.ECS.AddComponent(player, &components.Sprite{Image: playerImage.SelectSubImage(1, 1)})
	game.ECS.AddComponent(player, &components.Velocity{VX: 0, VY: 0, Acceleration: 5, Friction: 10})

	var velocity *components.Velocity
	if err := game.ECS.GetComponent(player, &velocity); err != nil {
		log.Fatal(err)
	}

	var sprite *components.Sprite
	if err := game.ECS.GetComponent(player, &sprite); err != nil {
		log.Fatal(err)
	}

	speed := 100.

	game.Input.RegisterHoldKeyEvent(
		cogni.KeyEvent{Key: ebiten.KeyA, Callback: func(delta float64) {
			velocity.DesiredVX = -speed
		}},
		cogni.KeyEvent{Key: ebiten.KeyD, Callback: func(delta float64) {
			velocity.DesiredVX = speed
		}},
		cogni.KeyEvent{Key: ebiten.KeyW, Callback: func(delta float64) {
			velocity.DesiredVY = -speed
		}},
		cogni.KeyEvent{Key: ebiten.KeyS, Callback: func(delta float64) {
			velocity.DesiredVY = speed
		}},
	)

	game.Input.RegisterReleaseKeyEvent(
		cogni.KeyEvent{Key: ebiten.KeyW, Callback: func(delta float64) {
			velocity.DesiredVY = 0
		}},
		cogni.KeyEvent{Key: ebiten.KeyA, Callback: func(delta float64) {
			velocity.DesiredVX = 0
		}},
		cogni.KeyEvent{Key: ebiten.KeyS, Callback: func(delta float64) {
			velocity.DesiredVY = 0
		}},
		cogni.KeyEvent{Key: ebiten.KeyD, Callback: func(delta float64) {
			velocity.DesiredVX = 0
		}},
	)

	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
