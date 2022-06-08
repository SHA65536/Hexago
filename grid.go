package Hexago

import (
	"fmt"
	"math"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

type Hexagon struct {
	// Fill Colours
	FR, FG, FB, FA float64
	//Stroke Colours and width
	SR, SG, SB, SA, SW float64
	//Font Colours font and text
	TR, TG, TB, TA float64
	Font           font.Face
	Text           string
	//Coords in the grid
	X, Y int
}

func (h *Hexagon) SetText(r, g, b, a float64, text string, font font.Face) {
	h.TR = r
	h.TG = g
	h.TB = b
	h.TA = a
	h.Text = text
	h.Font = font
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
	marginX       float64
	marginY       float64
}

// MakeHexGrid creates a <w> wide and <h> high grid with <r> rows and <c> columns
func MakeHexGrid(w, h, r, c float64) *HexGrid {
	rWidth := (2 * w) / ((3 * c) + 1)
	rHeight := (2 * h) / (math.Sqrt(3) * math.Sqrt((4*(r*r))+(4*r)+1))
	grid := &HexGrid{
		Context: gg.NewContext(int(w), int(h)),
		Tiles:   make([][]*Hexagon, int(r)),
		Width:   int(w),
		Height:  int(h),
		Rows:    int(r),
		Cols:    int(c),
		Radius:  math.Min(rWidth, rHeight),
	}
	if rHeight > rWidth {
		grid.marginY = h - ((0.5 + r) * (math.Sqrt(3 * rWidth * rWidth)))
		grid.marginY = grid.marginY / 2
	} else {
		grid.marginX = w - (((c / 2) * 3 * rHeight) + (rHeight / 2))
		grid.marginX = grid.marginX / 2
	}
	for i := range grid.Tiles {
		grid.Tiles[i] = make([]*Hexagon, int(c))
		for j := range grid.Tiles[i] {
			grid.Tiles[i][j] = &Hexagon{X: i, Y: j}
		}
	}
	return grid
}

// MakeHexGridWithContext creates a grid with external context
// Uses context width and height
func MakeHexGridWithContext(ctx *gg.Context, r, c float64) *HexGrid {
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	rWidth := (2 * w) / ((3 * c) + 1)
	rHeight := (2 * h) / (math.Sqrt(3) * math.Sqrt((4*(r*r))+(4*r)+1))
	grid := &HexGrid{
		Context: ctx,
		Tiles:   make([][]*Hexagon, int(r)),
		Width:   int(w),
		Height:  int(h),
		Rows:    int(r),
		Cols:    int(c),
		Radius:  math.Min(rWidth, rHeight),
	}
	if rHeight > rWidth {
		grid.marginY = h - ((0.5 + r) * (math.Sqrt(3 * rWidth * rWidth)))
		grid.marginY = grid.marginY / 2
	} else {
		grid.marginX = w - (((c / 2) * 3 * rHeight) + (rHeight / 2))
		grid.marginX = grid.marginX / 2
	}
	for i := range grid.Tiles {
		grid.Tiles[i] = make([]*Hexagon, int(c))
		for j := range grid.Tiles[i] {
			grid.Tiles[i][j] = &Hexagon{X: i, Y: j}
		}
	}
	return grid
}

// SetFillAll sets all of the tiles' fill colour
func (grid *HexGrid) SetFillAll(r, g, b, a float64) {
	for i := range grid.Tiles {
		for j := range grid.Tiles[i] {
			grid.Tiles[i][j].SetFill(r, g, b, a)
		}
	}
}

// SetStrokeAll sets all of the tiles' stroke colour and width
func (grid *HexGrid) SetStrokeAll(r, g, b, a, w float64) {
	for i := range grid.Tiles {
		for j := range grid.Tiles[i] {
			grid.Tiles[i][j].SetStroke(r, g, b, a, w)
		}
	}
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

// SetText sets the colour, size, and content of a cell's text
func (grid *HexGrid) SetText(x, y int, r, g, b, a float64, txt string, fontSize float64) error {
	if x >= grid.Rows || y >= grid.Cols {
		return fmt.Errorf("index error: index [%v][%v] out of bounds (%v,%v)", x, y, grid.Rows, grid.Cols)
	}
	font, _ := truetype.Parse(goregular.TTF)
	face := truetype.NewFace(font, &truetype.Options{Size: fontSize})
	grid.Tiles[x][y].SetText(r, g, b, a, txt, face)
	return nil
}

// DrawFunc recieves a function and applies that function to each cell of the grid
// the function recieves coordinates of each cell (x,y), the rows and cols (h,w)
// and a pointer to the cell itself
func (grid *HexGrid) DrawFunc(function func(x, y, h, w int, cell *Hexagon) error) error {
	for i := range grid.Tiles {
		for j := range grid.Tiles[i] {
			err := function(i, j, grid.Rows, grid.Cols, grid.Tiles[i][j])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// GetNeighbors returns the direct neighbors of a given cell
func (grid *HexGrid) GetNeighbors(x, y int) ([]*Hexagon, error) {
	neighbors := make([]*Hexagon, 0)
	if x >= grid.Rows || y >= grid.Cols {
		return nil, fmt.Errorf("index error: index [%v][%v] out of bounds (%v,%v)", x, y, grid.Rows, grid.Cols)
	}

	if x+1 < grid.Rows {
		// Checking above
		neighbors = append(neighbors, grid.Tiles[x+1][y])
	}
	if x > 0 {
		// ... and below
		neighbors = append(neighbors, grid.Tiles[x-1][y])
	}
	if y > 0 {
		// To the left
		neighbors = append(neighbors, grid.Tiles[x][y-1])
	}
	if y+1 < grid.Cols {
		// To the right
		neighbors = append(neighbors, grid.Tiles[x][y+1])
	}

	// Row specific cases
	if y%2 == 0 {
		if x+1 < grid.Rows && y > 0 {
			neighbors = append(neighbors, grid.Tiles[x+1][y-1])
		}
		if x+1 < grid.Rows && y+1 < grid.Cols {
			neighbors = append(neighbors, grid.Tiles[x+1][y+1])
		}
	} else {
		if x > 0 && y > 0 {
			neighbors = append(neighbors, grid.Tiles[x-1][y-1])
		}
		if x > 0 && y+1 < grid.Cols {
			neighbors = append(neighbors, grid.Tiles[x-1][y+1])
		}
	}
	return neighbors, nil
}

// SavePNG saves the grid as a PNG
func (grid *HexGrid) SavePNG(path string) error {
	grid.DrawGrid()
	return grid.Context.SavePNG(path)
}

func (grid *HexGrid) DrawGrid() {
	height := math.Sqrt(3 * grid.Radius * grid.Radius)
	for i := range grid.Tiles {
		for j := range grid.Tiles[i] {
			hex := grid.Tiles[i][j]
			x, y := grid.marginX, grid.marginY
			if j%2 == 0 {
				x += grid.Radius + (float64(j) * 1.5 * grid.Radius)
				y += height + (height * float64(i))
				grid.Context.DrawRegularPolygon(6, x, y, grid.Radius, 0)
			} else {
				x += grid.Radius + (float64(j) * 1.5 * grid.Radius)
				y += (height / 2) + (height * float64(i))
			}
			if hex.FA != 0 {
				grid.Context.DrawRegularPolygon(6, x, y, grid.Radius, 0)
				grid.Context.SetRGBA(hex.FR, hex.FG, hex.FB, hex.FA)
				grid.Context.Fill()
			}
			if hex.SW != 0 {
				grid.Context.DrawRegularPolygon(6, x, y, grid.Radius, 0)
				grid.Context.SetRGBA(hex.SR, hex.SG, hex.SB, hex.SA)
				grid.Context.SetLineWidth(hex.SW)
				grid.Context.Stroke()
			}
			if hex.Font != nil {
				grid.Context.SetFontFace(hex.Font)
				grid.Context.SetRGBA(hex.TR, hex.TG, hex.TB, hex.TA)
				grid.Context.DrawStringAnchored(hex.Text, x, y, 0.5, 0.5)
			}
		}
	}
}
