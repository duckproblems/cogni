package systems

import (
	"github.com/duckproblems/cogni/core/ecs"
	"github.com/duckproblems/cogni/core/ecs/components"
	"github.com/duckproblems/cogni/core/input"
	"github.com/hajimehoshi/ebiten/v2"
)

type PositionUpdateSystem struct{}

func (sys *PositionUpdateSystem) Update(ecsManager *ecs.ECSManager, inputManager *input.Manager, delta float64) {
	for _, entity := range ecsManager.Entities {
		var transform *components.Transform
		if entity.GetComponent(&transform) != nil {
			continue
		}

		var movement *components.Movement
		if entity.GetComponent(&movement) != nil {
			continue
		}

		transform.X += movement.VelocityX * delta
		transform.Y += movement.VelocityY * delta
	}
}

func (sys *PositionUpdateSystem) Draw(ecsManager *ecs.ECSManager, screen *ebiten.Image) {}
