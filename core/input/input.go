package input

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Manager struct {
	keyboard      [numKeys]keyboardState
	keyboardMutex sync.RWMutex
}

func New() *Manager {
	return &Manager{
		keyboard: [112]keyboardState{},
	}
}

func (m *Manager) Update() {
	m.keyboardMutex.Lock()
	defer m.keyboardMutex.Unlock()

	for i := 0; i < int(numKeys); i++ {
		key := Key(i).toEbiten()

		m.keyboard[i].JustPressed = inpututil.IsKeyJustPressed(key)
		m.keyboard[i].IsHolding = ebiten.IsKeyPressed(key)
		m.keyboard[i].JustReleased = inpututil.IsKeyJustReleased(key)
	}
}

func (m *Manager) KeyState(key Key) keyboardState {
	m.keyboardMutex.RLock()
	defer m.keyboardMutex.RUnlock()

	return m.keyboard[key]
}
