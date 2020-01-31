package main

//https://leetcode.com/problems/range-sum-query-2d-immutable/

type NumMatrix struct {
	sumRows [][]int
}

//First Iteration
// func Constructor(matrix [][]int) NumMatrix {
//     sumRows := make([][]int, len(matrix))
//     for i := range matrix {
//         sumRow := make([]int, len(matrix[i]))
//         sum := 0
//         for j := range matrix[i]{
//             sumRow[j] = sum + matrix[i][j]
//             sum = sumRow[j]
//         }
//         sumRows[i] = sumRow
//     }
//     return NumMatrix{
//         sumRows: sumRows,
//     }
// }

// func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
//     res := 0
//     for i:= row1; i <= row2; i++{
//         start := 0
//         if col1-1 >= 0 {
//             start = this.sumRows[i][col1-1]
//         }
//         res += this.sumRows[i][col2] - start
//     }
//     return res
// }

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{}
	}
	sumRows := make([][]int, len(matrix)+1)
	sumRows[0] = make([]int, len(matrix[0])+1)
	for i := 1; i < len(matrix)+1; i++ {
		sumRows[i] = make([]int, len(matrix[0])+1)
		for j := 1; j < len(matrix[0])+1; j++ {
			sumRows[i][j] = sumRows[i-1][j] + sumRows[i][j-1] + matrix[i-1][j-1] - sumRows[i-1][j-1]
		}
	}

	return NumMatrix{
		sumRows: sumRows,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if this.sumRows == nil {
		return 0
	}
	return this.sumRows[row2+1][col2+1] + this.sumRows[row1][col1] - this.sumRows[row1][col2+1] - this.sumRows[row2+1][col1]
}
