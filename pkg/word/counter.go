package word

import (
	"bufio"
	"os"
	"strings"
)

func Counter(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	wordCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		wordCount += len(words)
	}
	return wordCount, nil
}
