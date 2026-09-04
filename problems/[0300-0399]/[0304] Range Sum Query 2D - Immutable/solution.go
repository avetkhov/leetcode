package main

type NumMatrix struct {
	pref [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	pref := make([][]int, len(matrix)+1)
	for i := range pref {
		pref[i] = make([]int, len(matrix[0])+1)
	}

	for i, row := range matrix {
		for j := range row {
			pref[i+1][j+1] = matrix[i][j] + pref[i+1][j] + pref[i][j+1] - pref[i][j]
		}
	}

	return NumMatrix{pref: pref}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.pref[row2+1][col2+1] - this.pref[row2+1][col1] - this.pref[row1][col2+1] + this.pref[row1][col1]
}
