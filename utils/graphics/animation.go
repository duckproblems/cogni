package graphics

import (
	"math"

	"github.com/duckproblems/cogni/core/ecs/components"
	"github.com/hajimehoshi/ebiten/v2"
)

type AnimationState string

type AnimationSet map[AnimationState][]*ebiten.Image

const (
	AnimStateIdle      AnimationState = "Idle"
	AnimStateWalkUp                   = "WalkUp"
	AnimStateWalkDown                 = "WalkDown"
	AnimStateWalkLeft                 = "WalkLeft"
	AnimStateWalkRight                = "WalkRight"
)

func NewAnimationSet() AnimationSet {
	return make(AnimationSet)
}

func (as AnimationSet) AddAnimation(state AnimationState, frames []*ebiten.Image) {
	as[state] = frames
}

func (as AnimationSet) GetAnimationFrames(state AnimationState) []*ebiten.Image {
	return as[state]
}

func UpdateMovementAnimation(sprite *components.Sprite, animSet AnimationSet, vX, vY float64) {
	targetState := AnimStateIdle

	absX := math.Abs(vX)
	absY := math.Abs(vY)
	epsilon := 0.01

	if absX > epsilon || absY > epsilon {
		if absX >= absY {
			if vX < 0 {
				targetState = AnimStateWalkLeft
			} else {
				targetState = AnimStateWalkRight
			}
		} else {
			if vY < 0 {
				targetState = AnimStateWalkUp
			} else {
				targetState = AnimStateWalkDown
			}
		}
	}

	targetFrames, found := animSet[targetState]

	if !found || len(targetFrames) == 0 {
		sprite.Playing = false
		sprite.CurrentFrame = 0
		return
	}

	sprite.Playing = true
	sprite.Frames = targetFrames
}
