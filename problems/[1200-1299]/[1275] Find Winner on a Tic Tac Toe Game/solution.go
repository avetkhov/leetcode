package main

func tictactoe(moves [][]int) string {
	rows := [3]int{}
	cols := [3]int{}
	diag, antiDiag := 0, 0

	for i, move := range moves {
		r, c := move[0], move[1]
		val := 1

		if i%2 != 0 {
			val = -1
		}

		rows[r] += val
		cols[c] += val

		if r == c {
			diag += val
		}
		if r+c == 2 {
			antiDiag += val
		}

		if rows[r] == 3 || cols[c] == 3 || diag == 3 || antiDiag == 3 {
			return "A"
		}
		if rows[r] == -3 || cols[c] == -3 || diag == -3 || antiDiag == -3 {
			return "B"
		}
	}

	if len(moves) == 9 {
		return "Draw"
	}

	return "Pending"
}
