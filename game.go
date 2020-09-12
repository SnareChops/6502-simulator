package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/SnareChops/6502-simulator/visuals"
)

type Game struct{}

// Update updates a game by one tick. The given argument represents a screen image.
func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

// Draw draw the game screen. The given argument represents a screen image.
//
// (To be exact, Draw is not defined in this interface due to backward compatibility, but RunGame's
// behavior depends on the existence of Draw.)
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	screen.DrawImage(visuals.Image6502, visuals.Image6502Options)
}

// Layout accepts a native outside size in device-independent pixels and returns the game's logical
// screen size. On desktops, the outside is a window or a monitor (fullscreen mode)
//
// Even though the outside size and the screen size differ, the rendering scale is automatically
// adjusted to fit with the outside.
//
// You can return a fixed screen size if you don't care, or you can also return a calculated screen
// size adjusted with the given outside size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}
