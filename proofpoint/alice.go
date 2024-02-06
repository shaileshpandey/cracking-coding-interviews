package proofpoint

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func ReadWords(filePath string) []string {
	words := []string{}
	wordsWithFrequency := map[string]int{}
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Found Error while reading the file", err)
	} else {
		scannner := bufio.NewScanner(file)
		scannner.Split(bufio.ScanWords)

		for scannner.Scan() {
			word := scannner.Text()
			words = append(words, word)
			wordsWithFrequency[word] = wordsWithFrequency[word] + 1
		}
	}
	uniqueWords := getKeys(wordsWithFrequency)
	sort.Slice(uniqueWords, func(i, j int) bool {
		return wordsWithFrequency[uniqueWords[i]] > wordsWithFrequency[uniqueWords[j]]
	})
	DisplayTop10(uniqueWords)
	return uniqueWords
}

func DisplayTop10(words []string) {
	length := len(words)
	counter := 10
	if length < 10 {
		counter = length
	}
	for i := 0; i < counter; i++ {
		fmt.Println(words[i])
	}
}

func getKeys(dict map[string]int) []string {
	words := []string{}
	for key, _ := range dict {
		words = append(words, key)
	}
	return words
}
