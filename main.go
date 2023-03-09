package main

import (
	"github.com/ranon-rat/chess/lib/types"
)

func main() {
	//board := pieces.GenerateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
	//fmt.Println(board)
	//pieces.ShowBoard(board)
	//fmt.Println(pieces.CalculatePossibleMoves(1, 1, board))

	game := types.Game{}
	game.New()
}
