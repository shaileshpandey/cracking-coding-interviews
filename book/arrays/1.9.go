package arrays

import (
	"fmt"
	"strings"
)

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	input := s + s
	return strings.Index(input, goal) != -1
}

type MultiString struct {
	Input1 string
	Input2 string
	Output bool
}

func Call1_9() {
	dict := map[int]MultiString{
		1: {
			Input1: "abcde",
			Input2: "cdeab",
			Output: true,
		},

		2: {
			Input1: "abcde",
			Input2: "abced",
			Output: false,
		},
	}

	noissue := true
	for key, value := range dict {
		if rotateString(value.Input1, value.Input2) != value.Output {
			fmt.Printf("rotateString not returning result for %d", key)
			noissue = false
		}
	}

	if noissue {
		fmt.Println("No issues found in 1.9")
	}
}
