package main

import (
	"fmt"

	"github.com/SHA65536/Hexago"
)

type color struct {
	R float64
	G float64
	B float64
}

var colours = map[uint8]color{
	0: color{0.3, 0.3, 0.3},                              // Grey?
	1: color{51.0 / 255.0, 62.0 / 255.0, 212.0 / 255.0},  // Blue?
	2: color{47.0 / 255.0, 162.0 / 255.0, 54.0 / 255.0},  // Green?
	3: color{238.0 / 255.0, 222.0 / 255.0, 4.0 / 255.0},  // Yellow?
	4: color{247.0 / 255.0, 105.0 / 255.0, 21.0 / 255.0}, // Orange?
	5: color{253.0 / 255.0, 1.0 / 255.0, 0},              // Red?
}

func main() {
	// Create a 17*21 cell grid that is 800x800 pixels
	const rows, cols uint8 = 17, 21
	grid := Hexago.MakeHexGrid(800, 800, float64(rows), float64(cols))
	board := [rows][cols]uint8{}
	grid.SetStrokeAll(0, 0, 0, 1, 0.5)

	for x := 0; x < 100; x++ {
		board[rows/2][cols/2] += 1
		board[rows/2][cols/2] %= 6
		// Applying the  function for each cell
		grid.DrawFunc(func(i, j, _, _ int, _ *Hexago.Hexagon) error {
			color := colours[board[i][j]]
			return grid.SetFill(i, j, color.R, color.B, color.G, 1)
		})
		grid.SavePNG(fmt.Sprintf("output/%v.png", x))
	}
}
