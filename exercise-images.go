package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w int
	h int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel

}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

/* 各座標ごとの色情報を返す */
func (img Image) At(x, y int) color.Color {
	v := uint8(x * y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{w: 250, h: 250}
	pic.ShowImage(m)
}
