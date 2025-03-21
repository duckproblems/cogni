package main

import (
	"log"

	"github.com/duckproblems/cogni"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := cogni.New("Cogni Sample Window", 800, 600)

	game.Input.RegisterHoldKeyEvent(
		cogni.KeyEvent{Key: ebiten.KeyW, Callback: func() { println("Walk up") }},
		cogni.KeyEvent{Key: ebiten.KeyA, Callback: func() { println("Walk left") }},
		cogni.KeyEvent{Key: ebiten.KeyD, Callback: func() { println("Walk right") }},
		cogni.KeyEvent{Key: ebiten.KeyS, Callback: func() { println("Walk down") }},
	)

	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
