package main

import (
	"github.com/SHA65536/Hexago"
)

func main() {
	grid := Hexago.MakeHexGrid(1000, 1000, 6, 6)
	grid.SetFill(2, 2, 1, 0, 0, 1)
	grid.SetStroke(2, 2, 0, 0, 0, 1, 10)
	grid.SetFill(2, 3, 0, 1, 0, 1)
	grid.SetStroke(2, 3, 0, 0, 0, 1, 10)
	grid.SetFill(3, 3, 0, 0, 1, 1)
	grid.SetStroke(3, 3, 0, 0, 0, 1, 10)
	grid.SavePNG("out.png")
}
