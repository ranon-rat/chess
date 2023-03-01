package pieces

import (
	"fmt"
	"strings"
)

var piecesString = map[int]string{
	PAWN:   "p",
	KNIGHT: "n",
	BISHOP: "b",
	ROOK:   "r",
	QUEEN:  "q",
	KING:   "k",
}

func ShowBoard(board [8][8]*Piece) {
	for _, row := range board {
		for _, p := range row {
			if p == nil {
				fmt.Print(".")
				continue
			}
			str := piecesString[p.Kind]
			if !p.Black {
				str = strings.ToUpper(str)
			}
			fmt.Print(str)
		}
		fmt.Println()
	}
}
