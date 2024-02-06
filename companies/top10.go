package companies

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

func ReadFile() []string {
	wordsWithFrequency := map[string]int{}
	file, _ := os.Open("./companies/bigfile.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		wordsWithFrequency[word] = wordsWithFrequency[word] + 1
	}
	topWords := []string{}
	for key := range wordsWithFrequency {
		topWords = append(topWords, key)
	}

	sort.Slice(topWords, func(i, j int) bool {
		return wordsWithFrequency[topWords[i]] > wordsWithFrequency[topWords[j]]
	})

	return topWords
}
