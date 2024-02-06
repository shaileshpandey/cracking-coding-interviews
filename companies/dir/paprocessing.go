package dir

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

func findUniqueFromAddresses(rootDir string) (map[string]struct{}, error) {
	uniqueFromAddresses := make(map[string]struct{})
	var mutex sync.Mutex

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			go processFile(path, &mutex, uniqueFromAddresses)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return uniqueFromAddresses, nil
}

func processFile(filePath string, mutex *sync.Mutex, uniqueFromAddresses map[string]struct{}) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fromAddress := extractFromAddress(line)
		if fromAddress != "" {
			mutex.Lock()
			uniqueFromAddresses[fromAddress] = struct{}{}
			mutex.Unlock()
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
		return
	}
}

func extractFromAddress(line string) string {
	// Regular expression to extract FROM address
	re := regexp.MustCompile(`From:\s*<([^>]+)>`)
	match := re.FindStringSubmatch(line)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func main() {
	rootDir := "/path/to/emails/directory"

	uniqueFromAddresses, err := findUniqueFromAddresses(rootDir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Unique FROM addresses:")
	for address := range uniqueFromAddresses {
		fmt.Println(address)
	}
}
