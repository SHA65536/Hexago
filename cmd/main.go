package main

import (
	"math"

	"github.com/fogleman/gg"
)

func main() {
	rad := 100.0
	height := math.Sqrt(rad*rad + 2*rad*rad)
	dc := gg.NewContext(1000, 1000)
	dc.DrawRegularPolygon(6, 100, height, rad, 0)
	dc.SetRGB(1, 0, 0)
	dc.Fill()
	dc.DrawRegularPolygon(6, 250, height/2, rad, 0)
	dc.SetRGB(0, 1, 0)
	dc.Fill()
	dc.DrawRegularPolygon(6, 400, height, rad, 0)
	dc.SetRGB(0, 0, 1)
	dc.Fill()
	dc.SavePNG("out.png")
}
