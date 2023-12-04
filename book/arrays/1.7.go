package arrays

import (
	"fmt"
	"reflect"
)

func rotateMatrixExtraSpace(input [][]int) [][]int {
	length := len(input)
	output := [][]int{}
	for i := 0; i < length; i++ {
		var row []int
		for j := length - 1; j >= 0; j-- {
			row = append(row, input[j][i])
		}
		output = append(output, row)
	}
	return output
}

func rotateMatrixNoExtraSpace(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Step 2: Reverse each row of the transposed matrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0])/2; j++ {
			matrix[i][j], matrix[i][len(matrix[0])-j-1] = matrix[i][len(matrix[0])-j-1], matrix[i][j]
		}
	}
	return matrix
}

func Call1_7() {
	dict := map[int]TestCase{
		1: {
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			output: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},

		2: {
			input: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			output: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
	}

	noissue := true
	for key, value := range dict {
		if !reflect.DeepEqual(rotateMatrixExtraSpace(value.input), value.output) {
			fmt.Printf("rotateMatrixExtraSpace not returning result for %d", key)
			noissue = false
		}

		if !reflect.DeepEqual(rotateMatrixNoExtraSpace(value.input), value.output) {
			fmt.Printf("rotateMatrixExtraSpace not returning result for %d", key)
			noissue = false
		}
	}

	if noissue {
		fmt.Println("No issues found in 1.7")
	}
}

type TestCase struct {
	input  [][]int
	output [][]int
}
