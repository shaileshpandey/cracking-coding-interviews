package arrays

import (
	"fmt"
)

func onAway(first, second string) bool {
	if len(second) == len(first) {
		return oneEditReplace(first, second)
	} else if len(second) < len(first) {
		return oneEditInsert(second, first)
	}
	return oneEditInsert(first, second)
}

func oneEditReplace(s1, s2 string) bool {
	var foundDiff bool = false
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if foundDiff {
				return false
			}

			foundDiff = true
		}
	}
	return true
}

func oneEditInsert(s1, s2 string) bool {
	index1, index2 := 0, 0
	for index1 < len(s1) && index2 < len(s2) {
		if s1[index1] != s2[index2] {
			if index1 != index2 {
				return false
			}

			index2 = index2 + 1
		} else {
			index1++
			index2++
		}
	}
	return true
}

type input struct {
	first  string
	second string
}

func Call1_5() {
	dict := map[input]bool{
		{first: "pale", second: "ple"}:   true,
		{first: "pales", second: "pale"}: true,
		{first: "pale", second: "bale"}:  true,
		{first: "pale", second: "bake"}:  false,
	}

	noissue := true
	for key, value := range dict {
		if onAway(key.first, key.second) != value {
			fmt.Printf("onAway not returning result for %s", key)
			noissue = false
		}
	}

	if noissue {
		fmt.Println("No issues found in 1.5")
	}
}
