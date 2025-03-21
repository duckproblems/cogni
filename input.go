package cogni

import "github.com/hajimehoshi/ebiten/v2"

type inputManager struct {
	pressEvents   keyEventRegistry
	holdEvents    keyEventRegistry
	releaseEvents keyEventRegistry
}

type keyEventRegistry = map[ebiten.Key]keyEventHandler
type keyEventHandler = func()

type KeyEvent struct {
	Key      ebiten.Key
	Callback keyEventHandler
}

func newInputManager() *inputManager {
	return &inputManager{
		pressEvents:   keyEventRegistry{},
		holdEvents:    keyEventRegistry{},
		releaseEvents: keyEventRegistry{},
	}
}

func (i *inputManager) RegisterPressKeyEvent(events ...KeyEvent) {
	for _, event := range events {
		i.pressEvents[event.Key] = event.Callback
	}
}

func (i *inputManager) RegisterHoldKeyEvent(events ...KeyEvent) {
	for _, event := range events {
		i.holdEvents[event.Key] = event.Callback
	}
}

func (i *inputManager) RegisterReleaseKeyEvent(events ...KeyEvent) {
	for _, event := range events {
		i.releaseEvents[event.Key] = event.Callback
	}
}

func (i *inputManager) execPressEvent(key ebiten.Key) {
	event, found := i.pressEvents[key]
	if found {
		event()
	}
}

func (i *inputManager) execHoldEvent(key ebiten.Key) {
	event, found := i.holdEvents[key]
	if found {
		event()
	}
}

func (i *inputManager) execReleaseEvent(key ebiten.Key) {
	event, found := i.releaseEvents[key]
	if found {
		event()
	}
}
