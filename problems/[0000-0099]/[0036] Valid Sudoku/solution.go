package main

func isValidSudoku(board [][]byte) bool {
	size := len(board)
	boxSize := size / 3

	rows := make([]map[byte]struct{}, size)
	cols := make([]map[byte]struct{}, size)
	boxs := make([]map[byte]struct{}, size)

	for i := 0; i < size; i++ {
		rows[i] = make(map[byte]struct{})
		cols[i] = make(map[byte]struct{})
		boxs[i] = make(map[byte]struct{})
	}

	for i, row := range board {
		for j, ch := range row {
			if ch == '.' {
				continue
			}

			k := (i/boxSize)*boxSize + j/boxSize

			if _, ok := rows[i][ch]; ok {
				return false
			}
			if _, ok := cols[j][ch]; ok {
				return false
			}
			if _, ok := boxs[k][ch]; ok {
				return false
			}

			rows[i][ch], cols[j][ch], boxs[k][ch] = struct{}{}, struct{}{}, struct{}{}
		}
	}

	return true
}

func isValidSudoku(board [][]byte) bool {
	rows, cols, boxs := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}

	for i, row := range board {
		for j, ch := range row {
			if ch == '.' {
				continue
			}

			k := (i/3)*3 + j/3
			num := ch - '1'

			if rows[i][num] || cols[j][num] || boxs[k][num] {
				return false
			}

			rows[i][num], cols[j][num], boxs[k][num] = true, true, true
		}
	}

	return true
}
