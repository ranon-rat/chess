package pieces

const (
	KING   = 0
	QUEEN  = 1
	ROOK   = 2
	BISHOP = 3
	KNIGHT = 4
	PAWN   = 5
)

type Piece struct {
	Black   bool
	Moved   bool
	Passant bool
	Kind    int
	X       int
	Y       int
}
