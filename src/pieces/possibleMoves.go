package pieces

import (
	"errors"
)

// y,x
var (
	// this can move if a there is no piece
	BishopMoves = [][2]int{
		// up
		{1, 1}, {1, -1},
		//down
		{-1, -1}, {-1, 1}}
	// for now it will stay like this

	QueenMoves = [][2]int{
		// up
		{1, 1}, {1, -1}, {1, 0},
		//down
		{-1, -1}, {-1, 1}, {-1, 0},
		//left and right
		{0, -1}, {0, 1}}

	RookMoves = [][2]int{
		//right and left
		{0, 1}, {0, -1},
		//right and elft
		{-1, 0}, {1, 0}}

	KnightMoves = [][2]int{
		// weird L shapes
		{2, 1}, {-2, 1},
		{2, -1}, {-2, -1},
		{1, 2}, {-1, 2},
		{1, -2}, {-1, -2},
	}

	KingMoves = [][2]int{
		// up
		{1, 1}, {1, -1}, {1, 0},
		//down
		{-1, -1}, {-1, 1}, {-1, 0},
		//left and right
		{0, -1}, {0, 1},
	}

	Moves = map[int][][2]int{
		QUEEN:  QueenMoves,
		BISHOP: BishopMoves,
		ROOK:   RookMoves,
		KING:   KingMoves,
		KNIGHT: KnightMoves,
	}
	// pawns are differents so yeah
)

func CalculatePossibleMoves(x, y int, board [8][8]*Piece) (moves [][2]int, err error) {
	piece := board[y][x]

	if piece == nil {
		err = errors.New("in this square there is no piece")
		return
	}

	if piece.Kind >= QUEEN && piece.Kind <= BISHOP {
		for _, v := range Moves[piece.Kind] {
			moves = append(moves, CalculatePossibleLineMoves(x, y, piece, board, v)...)
		}
		return

	}

	if piece.Kind == PAWN {
		moves = CalculatePawnMoves(x, y, piece, board)
		return
	}
	for _, v := range Moves[piece.Kind] {
		moves = append(moves, CalculatePossibleSingleMoves(x, y, piece, board, v))
	}
	// castling
	if piece.Kind == KING && !piece.Moved {
		if !board[y][0].Moved {
			moves = append(moves, [2]int{y, x - 2})
		}
		if !board[y][7].Moved {
			moves = append(moves, [2]int{y, x + 2})
		}
	}

	return

}

func CalculatePossibleLineMoves(x, y int, piece *Piece, board [8][8]*Piece, move [2]int) (moves [][2]int) {
	dy, dx := move[0], move[1]
	x += dx
	y += dy

	for x >= 0 && y >= 0 && x < 8 && y < 8 {
		if board[y][x] != nil {
			if board[y][x].Black != piece.Black {
				moves = append(moves, [2]int{y, x})
			}
			break
		}
		moves = append(moves, [2]int{y, x})

		x += dx
		y += dy
	}
	return
}

func CalculatePossibleSingleMoves(x, y int, piece *Piece, board [8][8]*Piece, m [2]int) (move [2]int) {
	dy, dx := move[0], move[1]

	x += dx
	y += dy
	if x >= 0 && y >= 0 && x < 8 && y < 8 {
		if board[y][x] != nil {
			if board[y][x].Black == piece.Black {
				return

			}
		}
		move = [2]int{y, x}

	}
	return
}
