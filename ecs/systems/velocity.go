package systems

import (
	"github.com/duckproblems/cogni/ecs"
	"github.com/duckproblems/cogni/ecs/components"
	"github.com/hajimehoshi/ebiten/v2"
)

type Velocity struct{}

func (v Velocity) Update(manager *ecs.ECSManager, delta float64) {
	for entity := range manager.Components {
		var vel *components.Velocity
		if err := manager.GetComponent(entity, &vel); err != nil {
			continue
		}

		var trans *components.Transform
		if err := manager.GetComponent(entity, &trans); err != nil {
			continue
		}

		vel.VX += (vel.DesiredVX - vel.VX) * vel.Acceleration * delta
		vel.VY += (vel.DesiredVY - vel.VY) * vel.Acceleration * delta

		trans.X += vel.VX * delta
		trans.Y += vel.VY * delta

		if vel.DesiredVX == 0 && vel.DesiredVY == 0 {
			dampingFactor := 1.0 - (vel.Friction * delta)
			if dampingFactor < 0 {
				dampingFactor = 0
			}

			vel.VX *= dampingFactor
			vel.VY *= dampingFactor
		}
	}
}

func (v Velocity) Draw(manager *ecs.ECSManager, screen *ebiten.Image) {}
