package main

import (
	"fmt"

	"github.com/SHA65536/Hexago"
)

func main() {
	grid := Hexago.MakeHexGrid(1000, 1000, 6, 6)
	grid.SetFillAll(0, 0, 1, 1)
	grid.SetStrokeAll(0, 0, 0, 1, 10)
	for i := range grid.Tiles {
		for j := range grid.Tiles[i] {
			pos := fmt.Sprintf("%v,%v", i, j)
			grid.SetText(i, j, 0, 0, 0, 1, pos, 32)
		}
	}
	grid.SavePNG("out.png")
}
