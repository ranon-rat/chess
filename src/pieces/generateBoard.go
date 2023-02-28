package pieces

import (
	"strconv"
)

var pieces = map[rune]*Piece{
	'P': {Kind: PAWN, Black: false},
	'N': {Kind: KNIGHT, Black: false},
	'B': {Kind: BISHOP, Black: false},
	'R': {Kind: ROOK, Black: false},
	'Q': {Kind: QUEEN, Black: false},
	'S': {Kind: PAWN, Passant: true},
	'K': {Kind: KING, Black: false},
	'p': {Kind: PAWN, Black: true},
	'n': {Kind: KNIGHT, Black: true},
	's': {Kind: PAWN, Passant: true, Black: true},
	'b': {Kind: BISHOP, Black: true},
	'q': {Kind: QUEEN, Black: true},
	'r': {Kind: ROOK, Black: true},
	'k': {Kind: KING, Black: true},
}

func BoardToFen([][]*Piece) {

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
