package ecs

import (
	"fmt"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity int

type Component interface{}

type System interface {
	Update(manager *ECSManager, delta float64)
	Draw(manager *ECSManager, screen *ebiten.Image)
}

type ECSManager struct {
	nextEntity Entity
	systems    []System

	Components map[Entity]map[reflect.Type]Component
}

func New() *ECSManager {
	return &ECSManager{
		nextEntity: 1,
		systems:    []System{},
		Components: make(map[Entity]map[reflect.Type]Component),
	}
}

func (ecs *ECSManager) CreateEntity() Entity {
	entity := ecs.nextEntity
	ecs.nextEntity++

	return entity
}

func (ecs *ECSManager) AddSystems(systems ...System) {
	ecs.systems = systems
}

func (ecs *ECSManager) AddComponent(entity Entity, component Component) {
	if ecs.Components[entity] == nil {
		ecs.Components[entity] = make(map[reflect.Type]Component)
	}

	ecs.Components[entity][reflect.TypeOf(component)] = component
}

func (ecs *ECSManager) GetComponent(entity Entity, dst interface{}) error {
	comps, exists := ecs.Components[entity]
	if !exists {
		return fmt.Errorf("entity %d not found", entity)
	}

	t := reflect.TypeOf(dst).Elem()
	comp, exists := comps[t]
	if !exists {
		return fmt.Errorf("component of type %s not found for this entity %d", t, entity)
	}

	reflect.ValueOf(dst).Elem().Set(reflect.ValueOf(comp))
	return nil
}

func (ecs *ECSManager) Update(delta float64) {
	for _, system := range ecs.systems {
		system.Update(ecs, delta)
	}
}

func (ecs *ECSManager) Draw(screen *ebiten.Image) {
	for _, system := range ecs.systems {
		system.Draw(ecs, screen)
	}
}
