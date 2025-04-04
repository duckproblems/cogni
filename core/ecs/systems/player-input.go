package systems

import (
	"math"

	"github.com/duckproblems/cogni/core/ecs"
	"github.com/duckproblems/cogni/core/ecs/components"
	"github.com/duckproblems/cogni/core/input"
	"github.com/duckproblems/cogni/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerInputSystem struct{}

func (sys *PlayerInputSystem) Update(ecsManager *ecs.ECSManager, inputManager *input.Manager, delta float64) {
	var inputDx, inputDy float64
	if inputManager.KeyState(input.KB_W).IsHolding {
		inputDy -= 1.
	}
	if inputManager.KeyState(input.KB_S).IsHolding {
		inputDy += 1.
	}
	if inputManager.KeyState(input.KB_A).IsHolding {
		inputDx -= 1.
	}
	if inputManager.KeyState(input.KB_D).IsHolding {
		inputDx += 1.
	}

	distSq := utils.Sq(inputDx) + utils.Sq(inputDy)
	if distSq > utils.Epsilon {
		dist := math.Sqrt(distSq)
		inputDx /= dist
		inputDy /= dist
	} else {
		inputDx = 0
		inputDy = 0
	}

	for _, entity := range ecsManager.Entities {
		var input *components.InputControlled
		if entity.GetComponent(&input) != nil {
			continue
		}

		var movement *components.Movement
		if entity.GetComponent(&movement) != nil {
			continue
		}

		var intent *components.MovementIntent
		if entity.GetComponent(&intent) != nil {
			continue
		}

		intent.TargetX = inputDx
		intent.TargetY = inputDy
	}
}

func (sys *PlayerInputSystem) Draw(ecsManager *ecs.ECSManager, screen *ebiten.Image) {}
