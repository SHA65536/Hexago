package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/SHA65536/Hexago"
)

func main() {
	r, _ := strconv.Atoi(os.Args[1])
	c, _ := strconv.Atoi(os.Args[2])
	grid := Hexago.MakeHexGrid(1000, 1000, float64(r), float64(c))
	grid.SetFillAll(0, 0, 1, 1)
	grid.SetStrokeAll(0, 0, 0, 1, 10)
	grid.DrawFunc(func(i, j, _, _ int, _ *Hexago.Hexagon) error {
		pos := fmt.Sprintf("(%v,%v)", i, j)
		return grid.SetText(i, j, 0, 0, 0, 1, pos, 50)
	})
	neighbors, _ := grid.GetNeighbors(3, 2)
	for _, h := range neighbors {
		grid.SetFill(h.X, h.Y, 1, 0, 0, 1)
	}

	neighbors, _ = grid.GetNeighbors(1, 3)
	for _, h := range neighbors {
		grid.SetFill(h.X, h.Y, 0, 1, 0, 1)
	}
	grid.SavePNG("out.png")
}
