package main

import (
	"fmt"

	"github.com/ranon-rat/chess/src/pieces"
)

func main() {
	board := pieces.GenerateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
	pieces.ShowBoard(board)
	fmt.Println(pieces.CalculatePossibleMoves(1, 1, board))
}
