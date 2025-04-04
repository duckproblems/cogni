package systems

import (
	"github.com/duckproblems/cogni/core/ecs"
	"github.com/duckproblems/cogni/core/ecs/components"
	"github.com/duckproblems/cogni/core/input"
	"github.com/hajimehoshi/ebiten/v2"
)

type Animate struct{}

func (sys *Animate) Update(ecs *ecs.ECSManager, inputManager *input.Manager, delta float64) {
	for _, entity := range ecs.Entities {
		var sprite *components.Sprite
		if entity.GetComponent(&sprite) != nil {
			continue
		}

		if !sprite.Playing || len(sprite.Frames) <= 1 {
			continue
		}

		if sprite.CurrentFrame == 0 && sprite.FrameTimer == 0 && sprite.OnAnimationStart != nil {
			sprite.OnAnimationStart()
		}

		sprite.FrameTimer += delta
		frameDuration := 1.0 / sprite.FrameSpeed

		if sprite.FrameTimer >= frameDuration {
			sprite.FrameTimer = 0
			sprite.CurrentFrame++

			if sprite.OnAnimationFrame != nil {
				sprite.OnAnimationFrame()
			}

			if sprite.CurrentFrame >= len(sprite.Frames) {
				if sprite.Loop {
					sprite.CurrentFrame = 0
				} else {
					sprite.CurrentFrame = len(sprite.Frames) - 1
					sprite.Playing = false
				}

				if sprite.OnAnimationEnd != nil {
					sprite.OnAnimationEnd()
				}
			}
		}
	}
}

func (sys *Animate) Draw(ecs *ecs.ECSManager, screen *ebiten.Image) {
	for _, entity := range ecs.Entities {
		var sprite *components.Sprite
		if entity.GetComponent(&sprite) != nil {
			continue
		}

		var transform *components.Transform
		if entity.GetComponent(&transform) != nil {
			continue
		}

		if len(sprite.Frames) == 0 {
			continue
		}

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(transform.X, transform.Y)
		opts.GeoM.Rotate(transform.Rotation)
		opts.GeoM.Scale(transform.ScaleX, transform.ScaleY)

		if sprite.CurrentFrame >= len(sprite.Frames) {
			sprite.CurrentFrame = len(sprite.Frames) - 1
		}

		screen.DrawImage(sprite.Frames[sprite.CurrentFrame], opts)
	}
}
