# Hexago
A hexagonal grid package using [Go Graphics (gg)](https://github.com/fogleman/gg)

## Installation
    go get -u https://github.com/SHA65536/Hexago

## Usage
This is an example showing simple usage of the library:
```go
package main

import (
	"fmt"

	"github.com/SHA65536/Hexago"
)

func main() {
	// Create a 4*5 cell grid that is 800x600 pixels
	grid := Hexago.MakeHexGrid(800, 600, 4, 5)
	// Filling all cells with blue
	grid.SetFillAll(0, 0, 1, 1)
	// Giving all cells a 10 wide black border
	grid.SetStrokeAll(0, 0, 0, 1, 10)

	// Applying the  function for each cell
	grid.DrawFunc(func(i, j, _, _ int, _ *Hexago.Hexagon) error {
		// Writing the position on the cell
		pos := fmt.Sprintf("(%v,%v)", i, j)
		return grid.SetText(i, j, 0, 0, 0, 1, pos, 50)
	})

	// Getting all the neighbors of cell 3, 2
	neighbors, _ := grid.GetNeighbors(3, 2)
	for _, h := range neighbors {
		// Filling them with red
		grid.SetFill(h.X, h.Y, 1, 0, 0, 1)
	}

	// Getting all the neighbors of cell 1, 3
	neighbors, _ = grid.GetNeighbors(1, 3)
	for _, h := range neighbors {
		// Filling them with green
		grid.SetFill(h.X, h.Y, 0, 1, 0, 1)
	}

	// Saving grid to disk
	grid.SavePNG("out.png")
}
```
![Example](https://raw.githubusercontent.com/SHA65536/Hexago/main/cmd/example.png)