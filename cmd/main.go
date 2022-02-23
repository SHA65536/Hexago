package main

import (
	"github.com/SHA65536/Hexago"
)

func main() {
	grid := Hexago.MakeHexGrid(1000, 1000, 6, 6)
	grid.SetFillAll(0, 0, 1, 1)
	grid.SetStrokeAll(0, 0, 0, 1, 10)
	grid.SavePNG("out.png")
}
