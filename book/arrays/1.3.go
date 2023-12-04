package arrays

import (
	"fmt"
	"strings"
)

func urlfy(input string) string {
	input = strings.Trim(input, " ")
	return strings.ReplaceAll(input, " ", `%20`)
}

func Call1_3() {
	dict := map[string]string{
		"Mr. John Smith     ":       "Mr.%20John%20Smith",
		"Shailesh":                  "Shailesh",
		"Mr. Shailesh Kumar Pandey": "Mr.%20Shailesh%20Kumar%20Pandey",
	}

	for key, value := range dict {
		if urlfy(key) != value {
			fmt.Printf("urlfy not returning result for %s", key)
		}
	}

	fmt.Print("urlfy complete without any issue")
}
