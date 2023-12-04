package arrays

import (
	"fmt"
	"reflect"
)

func zeroMatrix(matrix [][]int) {
	rows_with_zero, columns_with_zeros := []int{}, []int{}

	rowLength := len(matrix)
	columnLength := len(matrix[0])
	for i := 0; i < rowLength; i++ {
		for j := 0; j < columnLength; j++ {
			if matrix[i][j] == 0 {
				rows_with_zero = append(rows_with_zero, i)
				columns_with_zeros = append(columns_with_zeros, j)
			}
		}
	}

	for i := 0; i < len(rows_with_zero); i++ {
		setRowCellsToZero(matrix, rows_with_zero[i])
	}

	for i := 0; i < len(columns_with_zeros); i++ {
		setColumnCellsToZero(matrix, columns_with_zeros[i])
	}
}

func setRowCellsToZero(matrix [][]int, rowNumber int) {
	columnLength := len(matrix[0])
	for i := 0; i < columnLength; i++ {
		matrix[rowNumber][i] = 0
	}
}

func setColumnCellsToZero(matrix [][]int, columnNumber int) {
	columnLength := len(matrix)
	for i := 0; i < columnLength; i++ {
		matrix[i][columnNumber] = 0
	}
}

func Call1_8() {
	dict := map[int]TestCase{
		1: {
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 0},
			},
			output: [][]int{
				{1, 2, 0},
				{4, 5, 0},
				{0, 0, 0},
			},
		},

		2: {
			input: [][]int{
				{5, 1, 9, 0},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{0, 14, 12, 16},
			},
			output: [][]int{
				{0, 0, 0, 0},
				{0, 4, 8, 0},
				{0, 3, 6, 0},
				{0, 0, 0, 0},
			},
		},
	}

	noissue := true
	for key, value := range dict {
		zeroMatrix(value.input)
		if !reflect.DeepEqual(value.input, value.output) {
			fmt.Printf("zeroMatrix not returning result for %d", key)
			noissue = false
		}
	}

	if noissue {
		fmt.Println("No issues found in 1.8")
	}
}
