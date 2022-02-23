package main

import (
	"github.com/SHA65536/Hexago"
)

func main() {
	grid := Hexago.MakeHexGrid(1000, 1000, 6, 6)
	grid.SetRGBA(2, 2, 1, 0, 0, 1)
	grid.SetRGBA(2, 3, 0, 1, 0, 1)
	grid.SetRGBA(3, 3, 0, 0, 1, 1)
	grid.SavePNG("out.png")
}
