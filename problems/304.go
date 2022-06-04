package problems

import "fmt"

func (p Problem) Run304() {
	matrix := [][]int{
		{3, 0, 1, 4, 2}, 
		{5, 6, 3, 2, 1}, 
		{1, 2, 0, 1, 5}, 
		{4, 1, 0, 1, 7}, 
		{1, 0, 3, 0, 5},
	}
	obj := Constructor(matrix)
	// Should be 11
	fmt.Println(obj.SumRegion(1, 1, 2, 2))
}

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	createSum := make([][]int, len(matrix) + 1)
	for index := range createSum {
		createSum[index] = make([]int, len(matrix[0]) + 1)
	}

	for i := range matrix {
		for j := range matrix[i] {
			createSum[i + 1][j + 1] = createSum[i + 1][j] + createSum[i][j + 1] + matrix[i][j] - createSum[i][j]
		}
	}
	return NumMatrix{sum: createSum}
}

func (matrix *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return matrix.sum[row2 + 1][col2 + 1] - matrix.sum[row1][col2+1]  - matrix.sum[row2+1][col1] + matrix.sum[row1][col1]
}
