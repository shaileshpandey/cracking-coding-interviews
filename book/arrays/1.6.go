package arrays

import (
	"fmt"
)

func stringCompression(input string) string {
	finalString := ""
	repeatedCount := 0
	for i := 0; i < len(input); i++ {
		repeatedCount = repeatedCount + 1
		if i+1 >= len(input) || input[i] != input[i+1] {
			finalString = fmt.Sprintf("%s%s%d", finalString, string(input[i]), repeatedCount)
			repeatedCount = 0
		}
	}

	if len(finalString) < len(input) {
		return finalString
	}
	return input
}

func Call1_6() {
	dict := map[string]string{
		"aabcccccaaa": "a2b1c5a3",
		"abcdef":      "abcdef",
	}

	noissue := true
	for key, value := range dict {
		if stringCompression(key) != value {
			fmt.Printf("stringCompression not returning result for %s", key)
			noissue = false
		}
	}

	if noissue {
		fmt.Println("No issues found in 1.6")
	}
}
