package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct{
	size int
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.size, i.size)
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) At(x, y int) color.Color {
	return color.RGBA{byte(x), byte(y), 0xff, 0xff}
}

func main() {
	m := &Image{255}
	pic.ShowImage(m)
}
