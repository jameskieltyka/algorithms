package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	var temp [][]rune
	for i := 0; i < n; i++ {
		temp = append(temp, []rune(strings.Repeat(".", n)))
	}
	canBePlaced(0, n, temp, &res)
	return res
}

func canBePlaced(col int, n int, temp [][]rune, res *[][]string) {
	if col == n {
		sol := make([]string, len(temp))
		for i := range sol {
			sol[i] = string(temp[i])
		}
		*res = append(*res, sol)
		return
	}

	for i := 0; i < n; i++ {
		temp[col][i] = 'Q'
		if !checkThreat(col, i, temp) {
			canBePlaced(col+1, n, temp, res)
		}
		temp[col][i] = '.'
	}
}

func checkThreat(col int, row int, temp [][]rune) bool {
	for i := 1; i <= col; i++ {
		if temp[col-i][row] == 'Q' {
			return true
		}

		if row+i < len(temp[0]) && temp[col-i][row+i] == 'Q' {
			return true
		}

		if row-i >= 0 && temp[col-i][row-i] == 'Q' {
			return true
		}
	}
	return false
}
