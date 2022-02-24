package main

import (
	"fmt"

	"github.com/SHA65536/Hexago"
)

const rows, cols int8 = 17, 21

type color struct {
	R float64
	G float64
	B float64
}

var colours = map[int8]color{
	0: {0.3, 0.3, 0.3},                              // Grey?
	1: {51.0 / 255.0, 62.0 / 255.0, 212.0 / 255.0},  // Blue?
	2: {47.0 / 255.0, 162.0 / 255.0, 54.0 / 255.0},  // Green?
	3: {238.0 / 255.0, 222.0 / 255.0, 4.0 / 255.0},  // Yellow?
	4: {247.0 / 255.0, 105.0 / 255.0, 21.0 / 255.0}, // Orange?
	5: {253.0 / 255.0, 1.0 / 255.0, 0},              // Red?
}

func main() {
	// Create a 17*21 cell grid that is 800x800 pixels
	grid := Hexago.MakeHexGrid(800, 800, float64(rows), float64(cols))
	board := [rows][cols]int8{}
	grid.SetStrokeAll(0, 0, 0, 1, 0.5)

	for x := 0; x < 1000; x++ {
		board[rows/2][cols/2] += 1
		solveBoard(grid, &board)
		// Applying the  function for each cell
		grid.DrawFunc(func(i, j, _, _ int, _ *Hexago.Hexagon) error {
			color := colours[board[i][j]]
			return grid.SetFill(i, j, color.R, color.B, color.G, 1)
		})
		grid.SavePNG(fmt.Sprintf("output/%v.png", x))
	}
}

func solveBoard(g *Hexago.HexGrid, b *[rows][cols]int8) {
	unsolved := true
	for unsolved {
		buffer := [rows][cols]int8{}
		unsolved = false
		g.DrawFunc(func(i, j, _, _ int, _ *Hexago.Hexagon) error {
			if b[i][j] > 5 {
				unsolved = true
				count := b[i][j] / 6
				buffer[i][j] += (-6 * count)
				around, _ := g.GetNeighbors(i, j)
				for _, cell := range around {
					buffer[cell.X][cell.Y] += count
				}
			}
			return nil
		})
		for i := range buffer {
			for j := range buffer[i] {
				b[i][j] += buffer[i][j]
			}
		}
	}
}
