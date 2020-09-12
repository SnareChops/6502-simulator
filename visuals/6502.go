package visuals

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/SnareChops/6502-simulator/utils"
)

var Image6502 = Init6502Image()
var Image6502Options = &ebiten.DrawImageOptions{
	GeoM:          ebiten.GeoM{},
	ColorM:        ebiten.ColorM{},
	CompositeMode: 0,
	Filter:        ebiten.FilterDefault,
}

func Init6502Image() *ebiten.Image {
	// Create new Image
	img := image.NewRGBA(image.Rect(0, 0, 428, 200))

	// Draw rectangle for processor body
	utils.Rect(img, image.Rect(0, 75, 427, 175), color.Black)

	// Draw pin rectangles
	for y := 65; y < 176; y += 110 {
		for x := 20; x < 401; x += 20 {
			utils.Rect(img, image.Rect(x, y, x+10, y+10), color.Black)
		}
	}

	// Convert to ebiten image
	result, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}

	Image6502Options.GeoM.Translate(426, 0)
	// return ebiten image
	return result
}
