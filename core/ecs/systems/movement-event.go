package systems

import (
	"github.com/duckproblems/cogni/core/ecs"
	"github.com/duckproblems/cogni/core/ecs/components"
	"github.com/duckproblems/cogni/core/input"
	"github.com/hajimehoshi/ebiten/v2"
)

type MovementEventSystem struct{}

func (sys *MovementEventSystem) Update(ecsManager *ecs.ECSManager, inputManager *input.Manager, delta float64) {
	for _, entity := range ecsManager.Entities {
		var movement *components.Movement
		if entity.GetComponent(&movement) != nil {
			continue
		}

		currentSpeed := movement.CurrentSpeed()
		isMoving := movement.IsCurrentlyMoving()

		if isMoving && !movement.WasMovingLastFrame {
			if movement.OnStartedMoving != nil {
				movement.OnStartedMoving(entity, movement.VelocityX, movement.VelocityY)
			}
		} else if !isMoving && movement.WasMovingLastFrame {
			if movement.OnStoppedMoving != nil {
				movement.OnStoppedMoving(entity, movement.VelocityX, movement.VelocityY)
			}
		}

		if isMoving {
			speedDiff := currentSpeed - movement.PreviousSpeed

			if speedDiff > components.SpeedChangeThreshold {
				if movement.OnAccelerating != nil {
					movement.OnAccelerating(entity, movement.VelocityX, movement.VelocityY)
				}
			} else if speedDiff < -components.SpeedChangeThreshold {
				if movement.OnDecelerating != nil {
					movement.OnDecelerating(entity, movement.VelocityX, movement.VelocityY)
				}
			} else {
				if movement.OnCruising != nil {
					movement.OnCruising(entity, movement.VelocityX, movement.VelocityY)
				}
			}
		}

		movement.WasMovingLastFrame = isMoving
		movement.PreviousSpeed = currentSpeed
	}
}

func (sys *MovementEventSystem) Draw(ecsManager *ecs.ECSManager, screen *ebiten.Image) {}
