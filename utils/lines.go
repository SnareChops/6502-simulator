package utils

import (
	"image"
	"image/color"
)

// HLine draws a horizontal line
func HLine(img *image.RGBA, x1, y, x2 int, col color.Color) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// VLine draws a veritcal line
func VLine(img *image.RGBA, x, y1, y2 int, col color.Color) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}

}

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(img *image.RGBA, rect image.Rectangle, col color.Color) {
	HLine(img, rect.Min.X, rect.Min.Y, rect.Max.X, col)
	HLine(img, rect.Min.X, rect.Max.Y, rect.Max.X, col)
	VLine(img, rect.Min.X, rect.Min.Y, rect.Max.Y, col)
	VLine(img, rect.Max.X, rect.Min.Y, rect.Max.Y, col)
}
