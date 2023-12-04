package arrays

import (
	"fmt"
	"slices"
	"strings"
)

func haveUniqueCharacters_BruteForce(str string) bool {
	length := len(str)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}

	return true
}

func haveUniqueCharacters_Sort(str string) bool {
	chars := strings.Split(str, "")
	slices.Sort(chars)
	for i := 0; i < len(chars); i++ {
		j := i + 1
		if j < len(chars) && chars[i] == chars[j] {
			return false
		}
	}
	return true
}

func Call1_1() {
	for _, input := range []string{"leetcode", "abc", "shailesh", "great"} {
		fmt.Println(fmt.Sprintf("'%s' has Unique characters", input), haveUniqueCharacters_BruteForce(input))
		fmt.Println(fmt.Sprintf("'%s' has Unique characters", input), haveUniqueCharacters_Sort(input))
	}
}
