package board

import piece "game/Piece"

const (
	BoardSize = 8
)

var (
	knightJumps = [][2]int{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {1, 2}, {1, -2}, {-1, 2}, {-1, -2}}
	rookDirs    = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	bishopDirs  = [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	allDirs     = append(append([][2]int{}, rookDirs...), bishopDirs...)
)

type Board [BoardSize][BoardSize]piece.Piece

type Square struct {
	Rank, File int
}

//Function to check if the value is in the Board
func InBound(r, f int) bool {
	return r >= 0 && r < BoardSize && f >= 0 && f < BoardSize
}

//Function are promissing return a pointer
func NewBoard() *Board {
	b := Board{}
	back := [8]piece.Piece{piece.WRook, piece.WKnight, piece.WBishop, piece.WQueen, piece.WKing, piece.WBishop, piece.WKnight, piece.WRook}

	for f := 0; f < 8; f++ {
		b[0][f] = back[f]     // White pieces on rank 1
		b[1][f] = piece.WPawn // White pawns on rank 2
		b[6][f] = piece.BPawn // Black pawns on rank 7
		b[7][f] = back[f] + 6 // Black pieces on rank 8

	}

	return &b
}
