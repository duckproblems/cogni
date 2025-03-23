package systems

import (
	"github.com/duckproblems/cogni/ecs"
	"github.com/duckproblems/cogni/ecs/components"
	"github.com/hajimehoshi/ebiten/v2"
)

type Render struct{}

func (r Render) Update(manager *ecs.ECSManager, delta float64) {}

func (r Render) Draw(manager *ecs.ECSManager, screen *ebiten.Image) {
	for entity := range manager.Components {
		var transform *components.Transform
		var sprite *components.Sprite

		if manager.GetComponent(entity, &transform) == nil {
			op := &ebiten.DrawImageOptions{}

			op.GeoM.Rotate(transform.Rotation)
			op.GeoM.Translate(transform.X, transform.Y)

			if manager.GetComponent(entity, &sprite) == nil {
				screen.DrawImage(sprite.Image, op)
			}
		}
	}
}
