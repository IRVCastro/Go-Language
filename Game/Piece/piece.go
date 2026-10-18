package piece

type Piece int

const (
	Empty   Piece = iota
	WPawn         // 1
	WKnight       // 2
	WBishop       // 3
	WRook         // 4
	WQueen        // 5
	WKing         // 6
	BPawn         // 7
	BKnight       // 8
	BBishop       // 9
	BRook         // 10
	BQueen        // 11
	BKing         // 12
)

// Check whether the number is between 1 and 6 which is the white range
func (p Piece) IsWhite() bool { return p >= WPawn && p <= WKing }

// Checks wheter the number is 7 or more, which is the balck range
func (p Piece) IsBlack() bool { return p >= BPawn }

// Kind returns the white version: BKnight and WKnight both return WKnight
func (p Piece) Kind() Piece {
	if p.IsBlack() {
		return p - 6
	}
	return p
}

//SameColor reports whether two pieces are on the same side
func (p Piece) SameColor(other Piece) bool {
	return (p.IsWhite() && other.Iswhite()) || (p.IsBlack() && other.IsBlack())
}

func (p Piece) Iswhite() bool {
	panic("unimplemented")
}
