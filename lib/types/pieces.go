package types

const FEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"

type Piece struct {
	Black bool
	Position
	Type byte
}

/*
func PawnMove(from, to *Piece) error {
	fx := from.Position[0]
	fy := from.Position[1]

	tx := to.Position[0]
	ty := to.Position[1]

	var double byte = 1
	var n byte = 1

	if fx == tx && fy == ty {
		return nil
	}

	if from.Black {
		double = 6
		n = -n
	}

	if fx == tx {
		if (fy == double && ty == fy+n*2) || (ty == fy+n) {
			return nil
		}
	}

	if fx == tx-1 || fx == tx+1 {
		if ty == fy+n {
			return nil
		}
	}

	return errors.New("invalid move")
}

func BishopMove(from, to *Piece) error {
	fx := from.Position[0]
	fy := from.Position[1]

	tx := to.Position[0]
	ty := to.Position[1]

	if fx == tx || fy == ty {
		return errors.New("invalid move")
	}

	dx := 1
	dy := 1

	if fx > tx && fy > ty {
		dx = -1
		dy = -1
	} else if fx > tx && fy < ty {
		dx = -1
		dy = 1
	} else if fx < tx && fy > ty {
		dx = 1
		dy = -1
	}

	x := fx + dx
	y := fy + dy

	for x != tx && y != ty {
		piece, err := board.GetPieceAt(Position{x, y})
		if err == nil && piece.Type != pieces.EMPTY {
			return errors.New("piece in the way")
		}
		x += dx
		y += dy
	}

	return nil
}

func RookMove(from, to *Piece) error {
	fx := from.Position[0]
	fy := from.Position[1]

	tx := to.Position[0]
	ty := to.Position[1]

	if fx == tx && fy == ty {
		return nil
	}

	if fx == tx && (fy > ty || fy < ty) {
		return nil
	}

	if fy == ty && (fx > tx || fx < tx) {
		return nil
	}

	return errors.New("invalid move")
}

func QueenMove(from, to *Piece) error {
	if err := BishopMove(from, to); err != nil {
		return err
	}

	if err := RookMove(from, to); err != nil {
		return err
	}

	return nil
}

func KingMove(from, to *Piece) error {
	fx := from.Position[0]
	fy := from.Position[1]

	tx := to.Position[0]
	ty := to.Position[1]

	if fx == tx && fy == ty {
		return nil
	}

	if fx == tx && (fy == ty-1 || fy == ty+1) {
		return nil
	}

	if fy == ty && (fx == tx-1 || fx == tx+1) {
		return nil
	}

	return nil
}

var isValidMove = map[byte]func(from, to *Piece) error{
	pieces.PAWN:   PawnMove,
	pieces.BISHOP: BishopMove,
	pieces.ROOK:   RookMove,
	pieces.QUEEN:  QueenMove,
	pieces.KING:   KingMove,
}
*/

var isValidMove = map[byte]func(from, to *Piece) error{}
