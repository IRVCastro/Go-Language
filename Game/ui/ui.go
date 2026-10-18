package ui

import (
	"embed"
	piece "game/Piece"
	"game/board"
	"image"
	"image/color"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

//go:embed images/*
var assets embed.FS
var PiecesSprites = map[piece.Piece]*ebiten.Image{
	piece.WPawn:   mustLoadImage("images/Wpawn.png"),
	piece.WRook:   mustLoadImage("images/Wrook.png"),
	piece.WKnight: mustLoadImage("images/Wknight.png"),
	piece.WBishop: mustLoadImage("images/Wbishop.png"),
	piece.WQueen:  mustLoadImage("images/Wqueen.png"),
	piece.WKing:   mustLoadImage("images/Wking.png"),
	piece.BPawn:   mustLoadImage("images/Bpawn.png"),
	piece.BRook:   mustLoadImage("images/Brook.png"),
	piece.BKnight: mustLoadImage("images/Bknight.png"),
	piece.BBishop: mustLoadImage("images/Bbishop.png"),
	piece.BQueen:  mustLoadImage("images/Bqueen.png"),
	piece.BKing:   mustLoadImage("images/Bking.png"),
}

type Game struct {
	Board *board.Board
}

const (
	BoardSize  = 8
	SquareSize = 80
)

func (g *Game) Update() error {
	return nil
}

// Function to display the layout
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 640
}

func (g *Game) Draw(screen *ebiten.Image) {

	//lineColor := color.RGBA{0, 0, 0, 255} // Black line color (Changed)

	// Define light and dark square colors
	lightColor := color.RGBA{242, 232, 231, 255}
	darkColor := color.RGBA{163, 82, 78, 255}

	//for lineDraw := 0; lineDraw <= 640; lineDraw += SquareSize {}

	// For To Create the colors of our Board
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			c := lightColor

			if (row+col)%2 == 1 {
				c = darkColor
			}
			//function to color the squares
			vector.FillRect(screen, float32(col*SquareSize), float32(row*SquareSize), SquareSize, SquareSize, c, false)
		}
	}

	// Rank Is Numbers on board
	// File Are characters on board
	for rank := 0; rank < BoardSize; rank++ {
		for file := 0; file < BoardSize; file++ {
			// Here i will get the actual position on board
			p := g.Board[rank][file]

			// If the position don't have anything i do nothing
			if p == piece.Empty {
				continue
			}

			sprite, ok := PiecesSprites[p]

			if !ok {
				continue
			}

			w := float64(sprite.Bounds().Dx())
			h := float64(sprite.Bounds().Dy())

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(SquareSize/w, SquareSize/h)
			op.GeoM.Translate(float64(file*SquareSize), float64((7-rank)*SquareSize))
			op.Filter = ebiten.FilterLinear
			screen.DrawImage(sprite, op)
		}
	}

}

func mustLoadImage(name string) *ebiten.Image {

	// Opens a file stream
	f, err := assets.Open(name)

	if err != nil {
		// Cause the program to crash immediately i f the file is missing
		panic(err)
	}

	// Cales the file stream to close automatically as soon as mustLoadImage Finish
	defer f.Close()

	//reads the raw byte stream from the file, detects its format (PNG), and decodes it into a standard Go image.Image interface in RAM.
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
