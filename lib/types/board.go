package types

import (
	"errors"
	"strconv"

	"github.com/ranon-rat/chess/lib/types/pieces"
)

var boardPieces = map[rune]*Piece{
	'P': {Type: pieces.PAWN},
	'N': {Type: pieces.KNIGHT},
	'B': {Type: pieces.BISHOP},
	'R': {Type: pieces.ROOK},
	'Q': {Type: pieces.QUEEN},
	'K': {Type: pieces.KING},
	'p': {Type: pieces.PAWN, Black: true},
	'n': {Type: pieces.KNIGHT, Black: true},
	'b': {Type: pieces.BISHOP, Black: true},
	'q': {Type: pieces.QUEEN, Black: true},
	'r': {Type: pieces.ROOK, Black: true},
	'k': {Type: pieces.KING, Black: true},
}

type Board struct {
	pieces []*Piece
}

func (b *Board) IsValidMove(from, to *Piece) error {
	return isValidMove[from.Type](from, to)
}

func (b *Board) OnEliminated(position Position) error {
	piece, err := b.GetPieceAt(position)

	if piece == nil || piece.Type == pieces.EMPTY {
		return errors.New("no piece at position")
	}

	if err != nil {
		return err
	}

	piece.Type = pieces.EMPTY

	return nil
}

func (b *Board) OnLimit(position Position) bool {
	if position[0] > 7 || position[1] > 7 {
		return true
	}

	return false
}

func (b *Board) GetPieceAt(position Position) (*Piece, error) {
	var err error = errors.New("no piece at position")

	for _, piece := range b.pieces {
		if piece.Position[0] == position[0] && piece.Position[1] == position[1] {
			return piece, nil
		}
	}

	return nil, err
}

func (b *Board) MovePiece(from, to Position) (err error) {
	if b.OnLimit(from) || b.OnLimit(to) {
		return errors.New("position out of bounds")
	}

	// - - -
	var pieceFrom, pieceTo *Piece

	if pieceFrom, err = b.GetPieceAt(from); err != nil {
		return err
	}

	if pieceTo, err = b.GetPieceAt(to); err != nil {
		return err
	}

	// - - -
	if pieceFrom.Type == pieces.EMPTY {
		return errors.New("no piece at position")
	}

	// - - -
	if err := b.IsValidMove(pieceFrom, pieceTo); err != nil {
		return err
	}

	// - - -
	pieceFrom.Position = pieceTo.Position
	pieceTo = &Piece{Type: pieces.EMPTY}

	return nil
}

func (b *Board) New() {
	var x byte = 0
	var y byte = 0

	for _, v := range FEN {
		if i, _ := strconv.Atoi(string(v)); i != 0 {
			x += byte(i - 1)

			continue
		}

		if v == '/' {
			x = 0
			y++

			continue
		}

		piece := *boardPieces[v]

		piece.Position[0] = x
		piece.Position[1] = y

		b.pieces = append(b.pieces, &piece)

		x++
	}
}
