package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (Image) At(x, y int) color.Color {
	ux := uint8(x)
	uy := uint8(y)
	return color.RGBA{ux*uy, ux^uy, (ux+uy)/2, 100}
}

func (Image) ColorModel() color.Model{
	return color.RGBAModel
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

