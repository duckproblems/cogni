package systems

import (
	"math"

	"github.com/duckproblems/cogni/core/ecs"
	"github.com/duckproblems/cogni/core/ecs/components"
	"github.com/hajimehoshi/ebiten/v2"
)

type Movement struct{}

func (sys Movement) Update(ecs *ecs.ECSManager, delta float64) {
	isUpPressed := ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp)
	isDownPressed := ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown)
	isLeftPressed := ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft)
	isRightPressed := ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight)

	for _, entity := range ecs.Entities {
		var transform *components.Transform
		if entity.GetComponent(&transform) != nil {
			continue
		}

		var input *components.Input
		if entity.GetComponent(&input) != nil {
			continue
		}

		var movement *components.Movement
		if entity.GetComponent(&movement) != nil {
			continue
		}

		var targetDx, targetDy float64
		if isUpPressed {
			targetDy -= 1
		}

		if isDownPressed {
			targetDy += 1
		}

		if isLeftPressed {
			targetDx -= 1
		}

		if isRightPressed {
			targetDx += 1
		}

		hasInput := targetDx != 0 || targetDy != 0
		if hasInput {
			dist := math.Sqrt(targetDx*targetDx + targetDy*targetDy)
			if dist > 1e-6 {
				targetDx /= dist
				targetDy /= dist
			}
		}

		if hasInput {
			accelAmount := movement.Acceleration * delta
			movement.VelocityX += targetDx * accelAmount
			movement.VelocityY += targetDy * accelAmount
		} else {
			movement.ApplyFriction(delta)
		}

		movement.ClampVelocity()

		currentSpeed := movement.CurrentSpeed()
		isMovingThisFrame := currentSpeed > components.VelocityStopThreshold

		if isMovingThisFrame && !movement.WasMovingLastFrame {
			if movement.OnStartedMoving != nil {
				movement.OnStartedMoving(movement.VelocityX, movement.VelocityY)
			}
		} else if !isMovingThisFrame && movement.WasMovingLastFrame {
			if movement.OnStoppedMoving != nil {
				movement.OnStoppedMoving(movement.VelocityX, movement.VelocityY)
			}
		}

		if isMovingThisFrame {
			speedDifference := currentSpeed - movement.PreviousSpeed

			if speedDifference > components.SpeedChangeThreshold {
				if movement.OnAccelerating != nil {
					movement.OnAccelerating(movement.VelocityX, movement.VelocityY)
				}
			} else if speedDifference < -components.SpeedChangeThreshold {
				if movement.OnDecelerating != nil {
					movement.OnDecelerating(movement.VelocityX, movement.VelocityY)
				}
			} else {
				if movement.OnCruising != nil {
					movement.OnCruising(movement.VelocityX, movement.VelocityY)
				}
			}
		}

		transform.X += movement.VelocityX * delta
		transform.Y += movement.VelocityY * delta

		movement.WasMovingLastFrame = isMovingThisFrame
		movement.PreviousSpeed = currentSpeed
	}
}

func (m Movement) Draw(ecs *ecs.ECSManager, screen *ebiten.Image) {}
