package pieces

const (
	KING   = iota
	QUEEN  = iota
	ROOK   = iota
	BISHOP = iota
	KNIGHT = iota
	PAWN   = iota
)

type Piece struct {
	Black   bool
	Moved   bool
	Passant bool
	Kind    int
	X       int
	Y       int
}
