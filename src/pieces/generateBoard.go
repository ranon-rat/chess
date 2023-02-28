package pieces

import (
	"strconv"
)

var pieces = map[rune]*Piece{
	'P': {Kind: PAWN, Passant: false, Moved: false, Black: false},
	'N': {Kind: KNIGHT, Passant: false, Moved: false, Black: false},
	'B': {Kind: BISHOP, Passant: false, Moved: false, Black: false},
	'R': {Kind: ROOK, Passant: false, Moved: false, Black: false},
	'Q': {Kind: QUEEN, Passant: false, Moved: false, Black: false},
	'K': {Kind: KING, Passant: false, Moved: false, Black: false},
	'p': {Kind: PAWN, Passant: false, Moved: false, Black: true},
	'n': {Kind: KNIGHT, Passant: false, Moved: false, Black: true},

	'b': {Kind: BISHOP, Passant: false, Moved: false, Black: true},
	'q': {Kind: QUEEN, Passant: false, Moved: false, Black: true},
	'r': {Kind: ROOK, Passant: false, Moved: false, Black: true},

	'k': {Kind: KING, Passant: false, Moved: false, Black: true},
}

func GenerateBoard(fen string) (board [][]*Piece) {
	board = make([][]*Piece, 8)
	board[0] = make([]*Piece, 8)
	x := 0
	y := 0
	for _, v := range fen {
		if val, _ := strconv.Atoi(string(v)); val != 0 {
			x += val - 1
			continue

		}
		if v == '/' {
			y++
			x = 0
			board[y] = make([]*Piece, 8)
			continue

		}
		board[y][x] = (pieces[v])
		x++

	}
	return
}
