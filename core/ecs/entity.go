package ecs

import (
	"fmt"
	"reflect"
)

type Entity struct {
	Name       string
	components map[reflect.Type]Component
}

func NewEntity(name string) Entity {
	return Entity{
		Name:       name,
		components: make(map[reflect.Type]Component),
	}
}

func (e *Entity) Rename(name string) {
	e.Name = name
}

func (e *Entity) AddComponent(component Component) {
	e.components[reflect.TypeOf(component)] = component
}

func (e *Entity) HasComponent(component Component) bool {
	_, exists := e.components[reflect.TypeOf(component)]
	return exists
}

func (e *Entity) GetComponent(target interface{}) error {
	t := reflect.TypeOf(target).Elem()
	comp, exists := e.components[t]
	if !exists {
		return fmt.Errorf("component of type %s not found for this entity %s", t, e.Name)
	}

	reflect.ValueOf(target).Elem().Set(reflect.ValueOf(comp))
	return nil
}
