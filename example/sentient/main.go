package main

import (
	"math"
	"slices"

	"github.com/duckproblems/cogni/core"
	"github.com/duckproblems/cogni/core/ecs"
	"github.com/duckproblems/cogni/core/ecs/components"
	"github.com/duckproblems/cogni/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := core.NewGame(*core.NewWindow(core.Window{Title: "Sentient", Geometry: utils.Pair[int]{Left: 800, Right: 600}}))

	playerImage, _ := utils.LoadFromPath("./assets/sprites/player.png", 16, 16)

	runningDown := []*ebiten.Image{
		playerImage.SelectSubImage(1, 1),
		playerImage.SelectSubImage(1, 2),
		playerImage.SelectSubImage(1, 3),
		playerImage.SelectSubImage(1, 4),
	}
	runningUp := []*ebiten.Image{
		playerImage.SelectSubImage(2, 1),
		playerImage.SelectSubImage(2, 2),
		playerImage.SelectSubImage(2, 3),
		playerImage.SelectSubImage(2, 4),
	}
	runningLeft := []*ebiten.Image{
		playerImage.SelectSubImage(3, 1),
		playerImage.SelectSubImage(3, 2),
		playerImage.SelectSubImage(3, 3),
		playerImage.SelectSubImage(3, 4),
	}
	runningRight := []*ebiten.Image{
		playerImage.SelectSubImage(4, 1),
		playerImage.SelectSubImage(4, 2),
		playerImage.SelectSubImage(4, 3),
		playerImage.SelectSubImage(4, 4),
	}

	player := ecs.NewEntity("Player")
	player.AddComponent(&components.Sprite{
		Frames:     runningDown,
		FrameSpeed: 5.,
		Loop:       true,
		Playing:    false,
	})
	player.AddComponent(&components.Transform{X: 30, Y: 30, ScaleX: 1, ScaleY: 1, Rotation: 0})

	var sprite *components.Sprite
	player.GetComponent(&sprite)

	updateAnimationDirection := func(vX, vY float64) {
		absX := math.Abs(vX)
		absY := math.Abs(vY)

		var targetFrames []*ebiten.Image

		if absX >= absY {
			if vX < 0 {
				targetFrames = runningLeft
			} else {
				targetFrames = runningRight
			}
		} else {
			if vY < 0 {
				targetFrames = runningUp
			} else {
				targetFrames = runningDown
			}
		}

		if !slices.Equal(sprite.Frames, targetFrames) && targetFrames != nil {
			sprite.Frames = targetFrames
			sprite.CurrentFrame = 0
			sprite.Playing = true
		}
	}

	player.AddComponent(&components.Input{})
	player.AddComponent(&components.Movement{
		MaxSpeed:     100.,
		Acceleration: 1000.,
		Friction:     0,
		OnStartedMoving: func(vX, vY float64) {
			sprite.Playing = true
			updateAnimationDirection(vX, vY)
		},
		OnStoppedMoving: func(vX, vY float64) {
			sprite.Playing = false
			sprite.CurrentFrame = 0
		},
		OnAccelerating: func(vX, vY float64) {
			updateAnimationDirection(vX, vY)
		},
		OnDecelerating: func(vX, vY float64) {
			updateAnimationDirection(vX, vY)
		},
		OnCruising: func(vX, vY float64) {
			updateAnimationDirection(vX, vY)
		},
	})

	game.ECS.AddEntity(player)
	game.Run()
}
