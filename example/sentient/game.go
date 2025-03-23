package main

import (
	"github.com/duckproblems/cogni/core"
	"github.com/duckproblems/cogni/utils"
)

func main() {
	game := core.NewGame(*core.NewWindow(core.Window{Title: "Sentient", Geometry: utils.Pair[int]{Left: 800, Right: 600}}))
	game.Run()
}
