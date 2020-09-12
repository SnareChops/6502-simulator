package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

const SCREEN_HEIGHT = 1280
const SCREEN_WIDTH = 720

func main() {
	game := &Game{}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(SCREEN_HEIGHT, SCREEN_WIDTH)
	ebiten.SetWindowTitle("6502 Simulator")

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
