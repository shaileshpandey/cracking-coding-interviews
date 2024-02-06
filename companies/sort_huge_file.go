package companies

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func externalMergeSort(inputFileName, outputFileName string, chunkSize int) error {
	// Open input file
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Read input file in chunks, sort each chunk, and write to temporary files
	chunkCount := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		chunk := make([]string, 0, chunkSize)
		for i := 0; i < chunkSize && scanner.Scan(); i++ {
			chunk = append(chunk, scanner.Text())
		}
		chunkCount++

		// Sort chunk
		sort.Strings(chunk)

		// Write sorted chunk to temporary file
		chunkFileName := fmt.Sprintf("chunk_%d.txt", chunkCount)
		chunkFile, err := os.Create(chunkFileName)
		if err != nil {
			return err
		}
		defer chunkFile.Close()

		for _, line := range chunk {
			fmt.Fprintln(chunkFile, line)
		}
	}

	// Merge sorted chunks
	var merge func(filenames []string) error
	merge = func(filenames []string) error {
		if len(filenames) == 1 {
			return nil
		}

		// Open input files
		files := make([]*os.File, 0, len(filenames))
		for _, filename := range filenames {
			file, err := os.Open(filename)
			if err != nil {
				return err
			}
			files = append(files, file)
			defer file.Close()
		}

		// Open output file for merged chunk
		mergedChunkFileName := fmt.Sprintf("merged_chunk_%d.txt", len(filenames))
		mergedChunkFile, err := os.Create(mergedChunkFileName)
		if err != nil {
			return err
		}
		defer mergedChunkFile.Close()

		// Merge sorted chunks
		scanners := make([]*bufio.Scanner, len(files))
		for i, file := range files {
			scanners[i] = bufio.NewScanner(file)
			scanners[i].Scan()
		}

		for {
			smallestIdx := -1
			smallestLine := ""

			// Find smallest line among current lines of all scanners
			for i, scanner := range scanners {
				if scanner.Err() != nil {
					continue
				}
				if smallestIdx == -1 || scanner.Text() < smallestLine {
					smallestIdx = i
					smallestLine = scanner.Text()
				}
			}

			// Exit loop if no smallest line found (all scanners EOF)
			if smallestIdx == -1 {
				break
			}

			// Write smallest line to merged chunk file
			fmt.Fprintln(mergedChunkFile, smallestLine)

			// Advance scanner to next line
			scanners[smallestIdx].Scan()
		}

		// Close input files
		for _, file := range files {
			file.Close()
		}

		// Recursively merge merged chunks
		return merge(append(filenames[:len(filenames)-1], mergedChunkFileName))
	}

	// Merge all sorted chunks
	err = merge(getChunkFileNames(chunkCount))
	if err != nil {
		return err
	}

	// Rename final merged chunk to output file
	os.Rename("merged_chunk_1.txt", outputFileName)

	return nil
}

func getChunkFileNames(n int) []string {
	filenames := make([]string, n)
	for i := range filenames {
		filenames[i] = fmt.Sprintf("chunk_%d.txt", i+1)
	}
	return filenames
}

func SortContents() {
	// Define input and output file names
	inputFileName := "large_dataset.txt"
	outputFileName := "sorted_dataset.txt"

	// Sort the large dataset using external merge sort
	err := externalMergeSort(inputFileName, outputFileName, 100000) // Chunk size: 100000 lines
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Large dataset sorted successfully.")
}
