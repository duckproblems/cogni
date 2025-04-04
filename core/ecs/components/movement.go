package components

import (
	"math"

	"github.com/duckproblems/cogni/core/ecs"
	"github.com/duckproblems/cogni/utils"
)

type MovementIntent struct {
	TargetX, TargetY float64
}

type Movement struct {
	VelocityX, VelocityY             float64
	MaxSpeed, Acceleration, Friction float64

	WasMovingLastFrame bool
	PreviousSpeed      float64

	OnStartedMoving func(entity ecs.Entity, vX, vY float64)
	OnStoppedMoving func(entity ecs.Entity, vX, vY float64)
	OnAccelerating  func(entity ecs.Entity, vX, vY float64)
	OnDecelerating  func(entity ecs.Entity, vX, vY float64)
	OnCruising      func(entity ecs.Entity, vX, vY float64)
}

const (
	VelocityStopThreshold float64 = 1.
	SpeedChangeThreshold  float64 = .1
)

func (m *Movement) ApplyFriction(delta float64) {
	if m.Friction <= 0 {
		return
	}

	if m.Friction >= 1 {
		m.VelocityX = 0
		m.VelocityY = 0
		return
	}

	frictionMultiplier := math.Pow(m.Friction, delta)
	m.VelocityX *= frictionMultiplier
	m.VelocityY *= frictionMultiplier

	if math.Abs(m.VelocityX) < utils.Epsilon {
		m.VelocityX = 0
	}
	if math.Abs(m.VelocityY) < utils.Epsilon {
		m.VelocityY = 0
	}
}

func (m *Movement) CurrentSpeed() float64 {
	return math.Sqrt(utils.Sq(m.VelocityX) + utils.Sq(m.VelocityY))
}

func (m *Movement) ClampVelocity() {
	speedSq := utils.Sq(m.VelocityX) + utils.Sq(m.VelocityY)
	maxSpeedSq := utils.Sq(m.MaxSpeed)

	if speedSq > maxSpeedSq && speedSq > utils.Epsilon {
		currentSpeed := math.Sqrt(speedSq)
		scale := m.MaxSpeed / currentSpeed

		m.VelocityX *= scale
		m.VelocityY *= scale
	}
}

func (m *Movement) IsCurrentlyMoving() bool {
	return (utils.Sq(m.VelocityX) + utils.Sq(m.VelocityY)) > utils.Sq(VelocityStopThreshold)
}
