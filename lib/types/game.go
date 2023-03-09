package types

import "fmt"

type Game struct {
	board Board
	turn  bool
}

func (g *Game) GetPieceAt(position Position) (Piece, error) {
	piece, err := g.board.GetPieceAt(position)

	return *piece, err
}

func (g *Game) MovePiece(from, to Position) error {
	return g.board.MovePiece(from, to)
}

func (g *Game) AlterTurn() {
	g.turn = !g.turn
}

func (g *Game) GetTurn() bool {
	return g.turn
}

func (g *Game) New() {
	g.AlterTurn()
	g.board.New()

	for _, piece := range g.board.pieces {
		fmt.Println(*piece)
	}
}
