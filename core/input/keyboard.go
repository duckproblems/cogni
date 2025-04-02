package input

import "github.com/hajimehoshi/ebiten/v2"

type keyboardState struct {
	JustPressed  bool
	IsHolding    bool
	JustReleased bool
}

type Key int

const (
	KB_A Key = iota
	KB_B
	KB_C
	KB_D
	KB_E
	KB_F
	KB_G
	KB_H
	KB_I
	KB_J
	KB_K
	KB_L
	KB_M
	KB_N
	KB_O
	KB_P
	KB_Q
	KB_R
	KB_S
	KB_T
	KB_U
	KB_V
	KB_W
	KB_X
	KB_Y
	KB_Z

	KB_0
	KB_1
	KB_2
	KB_3
	KB_4
	KB_5
	KB_6
	KB_7
	KB_8
	KB_9

	KB_F1
	KB_F2
	KB_F3
	KB_F4
	KB_F5
	KB_F6
	KB_F7
	KB_F8
	KB_F9
	KB_F10
	KB_F11
	KB_F12
	KB_F13
	KB_F14
	KB_F15
	KB_F16
	KB_F17
	KB_F18
	KB_F19
	KB_F20
	KB_F21
	KB_F22
	KB_F23
	KB_F24

	KB_ARROW_UP
	KB_ARROW_DOWN
	KB_ARROW_LEFT
	KB_ARROW_RIGHT

	KB_ESCAPE
	KB_SPACE
	KB_ENTER
	KB_BACKSPACE
	KB_CTRL
	KB_LCTRL
	KB_RCTRL
	KB_SHIFT
	KB_LSHIFT
	KB_RSHIFT
	KB_ALT
	KB_CAPSLOCK
	KB_TAB
	KB_QUOTE
	KB_OPEN_BRACKET
	KB_CLOSE_BRACKET
	KB_COMMA
	KB_DOT
	KB_BACKQUOTE
	KB_SEMICOLON
	KB_EQUALS
	KB_MINUS

	KB_PAUSE
	KB_SCROLLLOCK
	KB_PGUP
	KB_PGDOWN
	KB_DELETE
	KB_INSERT
	KB_HOME
	KB_END

	KB_FSLASH
	KB_BSLASH

	KB_NUMPAD_0
	KB_NUMPAD_1
	KB_NUMPAD_2
	KB_NUMPAD_3
	KB_NUMPAD_4
	KB_NUMPAD_5
	KB_NUMPAD_6
	KB_NUMPAD_7
	KB_NUMPAD_8
	KB_NUMPAD_9
	KB_NUMPAD_DIVIDE
	KB_NUMPAD_MULTIPLY
	KB_NUMPAD_SUBTRACT
	KB_NUMPAD_PLUS
	KB_NUMPAD_DOT
	KB_NUMPAD_ENTER

	numKeys
)

func (k Key) toEbiten() ebiten.Key {
	switch k {
	case KB_A:
		return ebiten.KeyA
	case KB_B:
		return ebiten.KeyB
	case KB_C:
		return ebiten.KeyC
	case KB_D:
		return ebiten.KeyD
	case KB_E:
		return ebiten.KeyE
	case KB_F:
		return ebiten.KeyF
	case KB_G:
		return ebiten.KeyG
	case KB_H:
		return ebiten.KeyH
	case KB_I:
		return ebiten.KeyI
	case KB_J:
		return ebiten.KeyJ
	case KB_K:
		return ebiten.KeyK
	case KB_L:
		return ebiten.KeyL
	case KB_M:
		return ebiten.KeyM
	case KB_N:
		return ebiten.KeyN
	case KB_O:
		return ebiten.KeyO
	case KB_P:
		return ebiten.KeyP
	case KB_Q:
		return ebiten.KeyQ
	case KB_R:
		return ebiten.KeyR
	case KB_S:
		return ebiten.KeyS
	case KB_T:
		return ebiten.KeyT
	case KB_U:
		return ebiten.KeyU
	case KB_V:
		return ebiten.KeyV
	case KB_W:
		return ebiten.KeyW
	case KB_X:
		return ebiten.KeyX
	case KB_Y:
		return ebiten.KeyY
	case KB_Z:
		return ebiten.KeyZ
	case KB_0:
		return ebiten.Key0
	case KB_1:
		return ebiten.Key1
	case KB_2:
		return ebiten.Key2
	case KB_3:
		return ebiten.Key3
	case KB_4:
		return ebiten.Key4
	case KB_5:
		return ebiten.Key5
	case KB_6:
		return ebiten.Key6
	case KB_7:
		return ebiten.Key7
	case KB_8:
		return ebiten.Key8
	case KB_9:
		return ebiten.Key9
	case KB_F1:
		return ebiten.KeyF1
	case KB_F2:
		return ebiten.KeyF2
	case KB_F3:
		return ebiten.KeyF3
	case KB_F4:
		return ebiten.KeyF4
	case KB_F5:
		return ebiten.KeyF5
	case KB_F6:
		return ebiten.KeyF6
	case KB_F7:
		return ebiten.KeyF7
	case KB_F8:
		return ebiten.KeyF8
	case KB_F9:
		return ebiten.KeyF9
	case KB_F10:
		return ebiten.KeyF10
	case KB_F11:
		return ebiten.KeyF11
	case KB_F12:
		return ebiten.KeyF12
	case KB_F13:
		return ebiten.KeyF13
	case KB_F14:
		return ebiten.KeyF14
	case KB_F15:
		return ebiten.KeyF15
	case KB_F16:
		return ebiten.KeyF16
	case KB_F17:
		return ebiten.KeyF17
	case KB_F18:
		return ebiten.KeyF18
	case KB_F19:
		return ebiten.KeyF19
	case KB_F20:
		return ebiten.KeyF20
	case KB_F21:
		return ebiten.KeyF21
	case KB_F22:
		return ebiten.KeyF22
	case KB_F23:
		return ebiten.KeyF23
	case KB_F24:
		return ebiten.KeyF24
	case KB_ARROW_UP:
		return ebiten.KeyArrowUp
	case KB_ARROW_DOWN:
		return ebiten.KeyArrowDown
	case KB_ARROW_LEFT:
		return ebiten.KeyArrowLeft
	case KB_ARROW_RIGHT:
		return ebiten.KeyArrowRight
	case KB_ESCAPE:
		return ebiten.KeyEscape
	case KB_SPACE:
		return ebiten.KeySpace
	case KB_ENTER:
		return ebiten.KeyEnter
	case KB_BACKSPACE:
		return ebiten.KeyBackspace
	case KB_CTRL:
		return ebiten.KeyControl
	case KB_LCTRL:
		return ebiten.KeyControlLeft
	case KB_RCTRL:
		return ebiten.KeyControlRight
	case KB_SHIFT:
		return ebiten.KeyShift
	case KB_LSHIFT:
		return ebiten.KeyShiftLeft
	case KB_RSHIFT:
		return ebiten.KeyShiftRight
	case KB_ALT:
		return ebiten.KeyAlt
	case KB_CAPSLOCK:
		return ebiten.KeyCapsLock
	case KB_TAB:
		return ebiten.KeyTab
	case KB_QUOTE:
		return ebiten.KeyQuote
	case KB_OPEN_BRACKET:
		return ebiten.KeyBracketLeft
	case KB_CLOSE_BRACKET:
		return ebiten.KeyBracketRight
	case KB_COMMA:
		return ebiten.KeyComma
	case KB_DOT:
		return ebiten.KeyPeriod
	case KB_BACKQUOTE:
		return ebiten.KeyBackquote
	case KB_SEMICOLON:
		return ebiten.KeySemicolon
	case KB_EQUALS:
		return ebiten.KeyEqual
	case KB_MINUS:
		return ebiten.KeyMinus
	case KB_PAUSE:
		return ebiten.KeyPause
	case KB_SCROLLLOCK:
		return ebiten.KeyScrollLock
	case KB_PGUP:
		return ebiten.KeyPageUp
	case KB_PGDOWN:
		return ebiten.KeyPageDown
	case KB_DELETE:
		return ebiten.KeyDelete
	case KB_INSERT:
		return ebiten.KeyInsert
	case KB_HOME:
		return ebiten.KeyHome
	case KB_END:
		return ebiten.KeyEnd
	case KB_FSLASH:
		return ebiten.KeySlash
	case KB_BSLASH:
		return ebiten.KeyBackslash
	case KB_NUMPAD_0:
		return ebiten.KeyNumpad0
	case KB_NUMPAD_1:
		return ebiten.KeyNumpad1
	case KB_NUMPAD_2:
		return ebiten.KeyNumpad2
	case KB_NUMPAD_3:
		return ebiten.KeyNumpad3
	case KB_NUMPAD_4:
		return ebiten.KeyNumpad4
	case KB_NUMPAD_5:
		return ebiten.KeyNumpad5
	case KB_NUMPAD_6:
		return ebiten.KeyNumpad6
	case KB_NUMPAD_7:
		return ebiten.KeyNumpad7
	case KB_NUMPAD_8:
		return ebiten.KeyNumpad8
	case KB_NUMPAD_9:
		return ebiten.KeyNumpad9
	case KB_NUMPAD_DIVIDE:
		return ebiten.KeyNumpadDivide
	case KB_NUMPAD_MULTIPLY:
		return ebiten.KeyNumpadMultiply
	case KB_NUMPAD_SUBTRACT:
		return ebiten.KeyNumpadSubtract
	case KB_NUMPAD_PLUS:
		return ebiten.KeyNumpadAdd
	case KB_NUMPAD_DOT:
		return ebiten.KeyNumpadDecimal
	case KB_NUMPAD_ENTER:
		return ebiten.KeyNumpadEnter
	default:
		return -1
	}
}

func (k Key) Name() string {
	return k.toEbiten().String()
}
