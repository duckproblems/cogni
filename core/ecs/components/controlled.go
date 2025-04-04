package components

import (
	"github.com/duckproblems/cogni/core/ecs"
	"github.com/duckproblems/cogni/core/ecs/ai"
)

type InputControlled struct{}

type AIControlled struct {
	TargetEntity ecs.Entity
	State        ai.State
}
