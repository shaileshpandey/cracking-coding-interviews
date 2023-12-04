package arrays

import (
	"fmt"
	"slices"
	"strings"
)

func isPermutation_Sort(str1 string, str2 string) bool {
	input1 := strings.Split(str1, "")
	input2 := strings.Split(str2, "")
	slices.Sort(input1)
	slices.Sort(input2)

	if len(input1) != len(input2) {
		return false
	}

	for i := 0; i < len(input1); i++ {
		if input1[i] != input2[i] {
			return false
		}
	}

	return true
}

func isPermutation_Map(str1 string, str2 string) bool {
	input1 := strings.Split(str1, "")
	input2 := strings.Split(str2, "")
	charMap := map[string]int{}

	for _, value := range input1 {
		charMap[value] = charMap[value] + 1
	}

	for _, value := range input2 {
		charMap[value] = charMap[value] - 1
	}

	for _, value := range charMap {
		if value != 0 {
			return false
		}
	}

	return true
}

func Call1_2() {
	dict := map[string]string{
		"ABC":      "CBA",
		"XYZ":      "ZYX",
		"Shailesh": "Akhilesh",
	}

	for key, value := range dict {
		fmt.Println(fmt.Sprintf("%s and %s are perm =>", key, value), isPermutation_Sort(key, value))
		fmt.Println(fmt.Sprintf("%s and %s are perm =>", key, value), isPermutation_Map(key, value))
	}
}
