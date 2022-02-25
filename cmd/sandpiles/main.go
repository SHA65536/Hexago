package main

import (
	"fmt"

	"github.com/SHA65536/Hexago"
	"github.com/schollz/progressbar/v3"
)

// Size of board
const rows, cols int8 = 35, 41

// Number of images to generate
const frames int = 3800

type color struct {
	R float64
	G float64
	B float64
}

var colours = map[int8]color{
	0: {0, 0, 0},             // Black
	1: {0.2, 0.243, 0.831},   // Blue
	2: {0.184, 0.635, 0.211}, // Green
	3: {1, 1, 0},             // Yellow
	4: {1, 0.549, 0},         // Orange
	5: {0.992, 0.01, 0},      // Red
}

func main() {
	// Create a 17*21 cell grid that is 1920x1080 pixels
	grid := Hexago.MakeHexGrid(1920, 1080, float64(rows), float64(cols))
	board := [rows][cols]int8{}

	// Creating progress bar
	bar := progressbar.Default(int64(frames))
	for x := 0; x < frames; x++ {
		bar.Add(1)

		// Adding a grain of sand to the center
		board[rows/2][cols/2] += 1
		solveBoard(grid, &board)

		// Applying the draw function for each cell
		grid.DrawFunc(func(i, j, _, _ int, _ *Hexago.Hexagon) error {
			color := colours[board[i][j]]
			return grid.SetFill(i, j, color.R, color.G, color.B, 1)
		})
		grid.SavePNG(fmt.Sprintf("output/%07d.png", x))
	}
}

// solveBoard takes in a sandpile board and topples the piles
// until the board is static
func solveBoard(g *Hexago.HexGrid, b *[rows][cols]int8) {
	unsolved := true
	for unsolved {
		// Buffer to save the sand that falls
		buffer := [rows][cols]int8{}

		unsolved = false
		g.DrawFunc(func(i, j, _, _ int, _ *Hexago.Hexagon) error {
			// If a cell has 6 or more grains, topple it
			if b[i][j] > 5 {
				unsolved = true
				count := b[i][j] / 6
				// The cell losses all it's sand except the division reminder
				buffer[i][j] += (-6 * count)
				around, _ := g.GetNeighbors(i, j)
				// Each neighbor gets 1/6 of the cell's sand
				for _, cell := range around {
					buffer[cell.X][cell.Y] += count
				}
			}
			return nil
		})
		// Applying buffer to the board
		for i := range buffer {
			for j := range buffer[i] {
				b[i][j] += buffer[i][j]
			}
		}
	}
}
