package components

import "math"

const (
	VelocityStopThreshold = 1.0
	SpeedChangeThreshold  = 0.1
)

type Movement struct {
	VelocityX, VelocityY             float64
	MaxSpeed, Acceleration, Friction float64

	OnStartedMoving func(vX, vY float64)
	OnStoppedMoving func(vX, vY float64)
	OnAccelerating  func(vX, vY float64)
	OnDecelerating  func(vX, vY float64)
	OnCruising      func(vX, vY float64)

	WasMovingLastFrame bool
	PreviousSpeed      float64
}

func (m *Movement) ApplyFriction(delta float64) {
	frictionMultiplier := math.Pow(m.Friction, delta)
	m.VelocityX *= frictionMultiplier
	m.VelocityY *= frictionMultiplier

	const epsilon = 1e-3
	if math.Abs(m.VelocityX) < epsilon {
		m.VelocityX = 0
	}
	if math.Abs(m.VelocityY) < epsilon {
		m.VelocityY = 0
	}
}

func (m *Movement) ClampVelocity() {
	speedSq := m.VelocityX*m.VelocityX + m.VelocityY*m.VelocityY
	maxSpeedSq := m.MaxSpeed * m.MaxSpeed

	if speedSq > maxSpeedSq && speedSq > 1e-6 {
		currentSpeed := math.Sqrt(speedSq)
		scale := m.MaxSpeed / currentSpeed
		m.VelocityX *= scale
		m.VelocityY *= scale
	}
}

func (m *Movement) CurrentSpeed() float64 {
	return math.Sqrt(m.VelocityX*m.VelocityX + m.VelocityY*m.VelocityY)
}

func (m *Movement) IsCurrentlyMoving() bool {
	return (m.VelocityX*m.VelocityX + m.VelocityY*m.VelocityY) > VelocityStopThreshold*VelocityStopThreshold
}
