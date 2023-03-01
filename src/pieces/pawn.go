package pieces

import "fmt"

// pawn are kinda weird so yeah, this is needed
func CalculatePawnMoves(x, y int, piece *Piece, board [8][8]*Piece) (moves [][2]int) {
	direction := 1
	if !piece.Black {
		direction = -1
	}
	attack := []int{-1, 1}
	// attacks
	for _, v := range attack {
		moves = append(moves, NormalAttack(x, y, v, direction, piece, board)...)
		// passant
		moves = append(moves, Passant(x, y, v, direction, piece, board)...)
	}
	if y+direction < 8 && y+direction >= 0 {
		if board[y+direction][x] == nil {
			moves = append(moves, [2]int{y + direction, x})
			//jump 2
			if board[y+direction*2][x] == nil && !piece.Moved {
				moves = append(moves, [2]int{y + direction*2, x})
			}
		}
		fmt.Println(moves)

	}
	fmt.Println(moves)

	return
}
func NormalAttack(x, y, attack, direction int, piece *Piece, board [8][8]*Piece) (moves [][2]int) {
	if board[y+direction][x+attack] == nil {
		return
	}
	if y+direction < 8 && y+direction >= 0 && x+attack >= 0 && x+attack < 8 {
		if board[y+direction][x+attack].Black != piece.Black {
			moves = append(moves, [2]int{(y + direction), x + attack})
		}

	}
	return

}
func Passant(x, y, attack, direction int, piece *Piece, board [8][8]*Piece) (moves [][2]int) {
	if board[y][x+attack] == nil {
		return
	}

	if x+attack >= 0 && x+attack < 8 {
		if board[y][x+attack].Kind == PAWN && board[y][x+attack].Passant && board[y][x+attack].Black != piece.Black {
			moves = append(moves, [2]int{(y + direction), x + attack})
		}

	}
	return
}
