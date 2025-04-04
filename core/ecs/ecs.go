package ecs

import (
	"github.com/duckproblems/cogni/core/input"
	"github.com/hajimehoshi/ebiten/v2"
)

type ECSManager struct {
	nextEntityID int

	Systems  []System
	Entities map[int]Entity
}

func New() *ECSManager {
	return &ECSManager{
		nextEntityID: 1,
		Systems:      make([]System, 0),
		Entities:     make(map[int]Entity),
	}
}

func (e *ECSManager) AddEntity(entity Entity) {
	e.Entities[e.nextEntityID] = entity
	e.nextEntityID++
}

func (e *ECSManager) AddSystem(system System) {
	e.Systems = append(e.Systems, system)
}

func (e *ECSManager) Update(inputManager *input.Manager, delta float64) {
	for _, system := range e.Systems {
		system.Update(e, inputManager, delta)
	}
}

func (e *ECSManager) Draw(screen *ebiten.Image) {
	for _, system := range e.Systems {
		system.Draw(e, screen)
	}
}
