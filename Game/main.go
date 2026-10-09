package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

// Update proceeds the game state
// Update is called every tick (1/60 [s] by default)
func (g *Game) Update() error {
	return nil
}

// Draw draw the game screen
// Draw is called every frame (typically 1/60[s] for 60Hz display)
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

// Layout takes the outside size and returns the logical screen size
// If you don't have to adjuts the screen size with the outside
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{}
	// Specify the window size as you like, here a doubled size is specified
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")
	// Cal ebiten.Run game to start your game loop

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
