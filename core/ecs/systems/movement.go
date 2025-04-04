package systems

import (
	"github.com/duckproblems/cogni/core/ecs"
	"github.com/duckproblems/cogni/core/ecs/components"
	"github.com/duckproblems/cogni/core/input"
	"github.com/hajimehoshi/ebiten/v2"
)

type MovementSystem struct{}

func (sys *MovementSystem) Update(ecsManager *ecs.ECSManager, inputManager *input.Manager, delta float64) {
	for _, entity := range ecsManager.Entities {
		var intent *components.MovementIntent
		if entity.GetComponent(&intent) != nil {
			continue
		}

		var movement *components.Movement
		if entity.GetComponent(&movement) != nil {
			continue
		}

		hasIntent := intent.TargetX != 0 || intent.TargetY != 0

		if hasIntent {
			accelAmount := movement.Acceleration * delta

			movement.VelocityX += intent.TargetX * accelAmount
			movement.VelocityY += intent.TargetY * accelAmount
		} else {
			movement.ApplyFriction(delta)
		}

		movement.ClampVelocity()

		intent.TargetX = 0
		intent.TargetY = 0
	}
}

func (sys *MovementSystem) Draw(ecsManager *ecs.ECSManager, screen *ebiten.Image) {}
