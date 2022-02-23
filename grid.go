package Hexago

import (
	"fmt"
	"math"

	"github.com/fogleman/gg"
)

type Hexagon struct {
	R, G, B, A float64
	X, Y       int
}

func (h *Hexagon) SetRGBA(r, g, b, a float64) {
	h.R = r
	h.G = g
	h.B = b
	h.A = a
}

type HexGrid struct {
	Context       *gg.Context
	Tiles         [][]*Hexagon
	Width, Height int
	Rows, Cols    int
	Radius        float64
}

func MakeHexGrid(w, h, r, c int) *HexGrid {
	grid := &HexGrid{
		Context: gg.NewContext(w, h),
		Tiles:   make([][]*Hexagon, r),
		Width:   w,
		Height:  h,
		Rows:    r,
		Cols:    c,
		Radius:  100,
	}
	for i := range grid.Tiles {
		grid.Tiles[i] = make([]*Hexagon, c)
		for j := range grid.Tiles[i] {
			grid.Tiles[i][j] = &Hexagon{X: i, Y: j}
		}
	}
	return grid
}

func (grid *HexGrid) SetRGBA(x, y int, r, g, b, a float64) error {
	if x >= grid.Rows || y >= grid.Cols {
		return fmt.Errorf("index error: index [%v][%v] out of bounds (%v,%v)", x, y, grid.Rows, grid.Cols)
	}
	grid.Tiles[x][y].SetRGBA(r, g, b, a)
	return nil
}

func (grid *HexGrid) SavePNG(path string) error {
	height := math.Sqrt(3 * grid.Radius * grid.Radius)
	for i := range grid.Tiles {
		for j := range grid.Tiles[i] {
			hex := grid.Tiles[i][j]
			if j%2 == 0 {
				x := grid.Radius + (float64(j) * 1.5 * grid.Radius)
				y := height + (height * float64(i))
				grid.Context.DrawRegularPolygon(6, x, y, grid.Radius, 0)
			} else {
				x := grid.Radius + (float64(j) * 1.5 * grid.Radius)
				y := (height / 2) + (height * float64(i))
				grid.Context.DrawRegularPolygon(6, x, y, grid.Radius, 0)
			}
			grid.Context.SetRGBA(hex.R, hex.G, hex.B, hex.A)
			grid.Context.Fill()
		}
	}
	return grid.Context.SavePNG(path)
}
