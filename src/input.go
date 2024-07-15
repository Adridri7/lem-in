package src

import (
	"os"
	"strings"
)

// Convert .txt file into lines.
func FileToTab(fileName string) []string {
	data, _ := os.ReadFile("./examples/" + fileName)
	return strings.Split(string(data), "\n")
}
