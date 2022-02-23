package Hexago

import (
	"fmt"
	"math"

	"github.com/fogleman/gg"
)

type Hexagon struct {
	FR, FG, FB, FA     float64
	SR, SG, SB, SA, SW float64
	X, Y               int
}

func (h *Hexagon) SetFill(r, g, b, a float64) {
	h.FR = r
	h.FG = g
	h.FB = b
	h.FA = a
}

func (h *Hexagon) SetStroke(r, g, b, a, w float64) {
	h.SR = r
	h.SG = g
	h.SB = b
	h.SA = a
	h.SW = w
}

type HexGrid struct {
	Context       *gg.Context
	Tiles         [][]*Hexagon
	Width, Height int
	Rows, Cols    int
	Radius        float64
}

// MakeHexGrid creates a <w> wide and <h> high grid with <r> rows and <c> columns
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

// SetFill sets the fill of a cell
func (grid *HexGrid) SetFill(x, y int, r, g, b, a float64) error {
	if x >= grid.Rows || y >= grid.Cols {
		return fmt.Errorf("index error: index [%v][%v] out of bounds (%v,%v)", x, y, grid.Rows, grid.Cols)
	}
	grid.Tiles[x][y].SetFill(r, g, b, a)
	return nil
}

// SetStroke sets the stroke of a cell
func (grid *HexGrid) SetStroke(x, y int, r, g, b, a, w float64) error {
	if x >= grid.Rows || y >= grid.Cols {
		return fmt.Errorf("index error: index [%v][%v] out of bounds (%v,%v)", x, y, grid.Rows, grid.Cols)
	}
	grid.Tiles[x][y].SetStroke(r, g, b, a, w)
	return nil
}

// SavePNG saves the grid as a PNG
func (grid *HexGrid) SavePNG(path string) error {
	height := math.Sqrt(3 * grid.Radius * grid.Radius)
	for i := range grid.Tiles {
		for j := range grid.Tiles[i] {
			hex := grid.Tiles[i][j]
			var x, y float64
			if j%2 == 0 {
				x = grid.Radius + (float64(j) * 1.5 * grid.Radius)
				y = height + (height * float64(i))
				grid.Context.DrawRegularPolygon(6, x, y, grid.Radius, 0)
			} else {
				x = grid.Radius + (float64(j) * 1.5 * grid.Radius)
				y = (height / 2) + (height * float64(i))
			}
			grid.Context.DrawRegularPolygon(6, x, y, grid.Radius, 0)
			grid.Context.SetRGBA(hex.FR, hex.FG, hex.FB, hex.FA)
			grid.Context.Fill()
			grid.Context.DrawRegularPolygon(6, x, y, grid.Radius, 0)
			grid.Context.SetRGBA(hex.SR, hex.SG, hex.SB, hex.SA)
			grid.Context.SetLineWidth(hex.SW)
			grid.Context.Stroke()
		}
	}
	return grid.Context.SavePNG(path)
}
