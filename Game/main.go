package main

import (
	"game/board"
	"game/ui"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	game := &ui.Game{
		Board: board.NewBoard(),
	}

	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Chess Game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

	//boardVar := board.InitBoard()
	//board.PrintBoard(boardVar)
}
