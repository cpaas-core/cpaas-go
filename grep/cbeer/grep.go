package grep

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Search a file for lines matching a regular expression pattern.
func Search(pattern string, flags, files []string) []string {
	searchResults := make([]string, 0)
	var printLineNumbers bool
	var printOnlyNames bool
	var includeFileNames bool

	for _, flagValue := range flags {
		switch flagValue {
		case "-n":
			printLineNumbers = true
		case "-l":
			printOnlyNames = true
		}
	}

	if len(files) > 1 {
		includeFileNames = true
	}

	for _, fileName := range files {
		fileHandle, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer fileHandle.Close()

		scanner := bufio.NewScanner(fileHandle)
		lineNumber := 1
		var lineOfText string

		if printOnlyNames {
			matchFound := false
			for scanner.Scan() {
				if isLineMatch(pattern, scanner.Text(), flags) {
					matchFound = true
				}
			}
			if matchFound {
				searchResults = append(searchResults, fileName)
			}
		} else {
			for scanner.Scan() {
				if isLineMatch(pattern, scanner.Text(), flags) {
					if includeFileNames {
						if printLineNumbers {
							lineOfText = fmt.Sprintf("%s:%d:%s", fileName, lineNumber, scanner.Text())
						} else {
							lineOfText = fmt.Sprintf("%s:%s", fileName, scanner.Text())
						}
					} else {
						if printLineNumbers {
							lineOfText = fmt.Sprintf("%d:%s", lineNumber, scanner.Text())
						} else {
							lineOfText = scanner.Text()
						}
					}
					searchResults = append(searchResults, lineOfText)
				}
				lineNumber++
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	return searchResults
}

func isLineMatch(pattern, line string, flags []string) bool {
	var caseInsensitive bool
	var invertProgram bool
	var matchEntireLines bool
	var foundMatch bool

	for _, flagValue := range flags {
		switch flagValue {
		case "-i":
			caseInsensitive = true
		case "-v":
			invertProgram = true
		case "-x":
			matchEntireLines = true
		}
	}

	if matchEntireLines {
		if caseInsensitive {
			foundMatch = strings.EqualFold(line, pattern)
		} else {
			foundMatch = pattern == line
		}
	} else {
		if caseInsensitive {
			foundMatch = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
		} else {
			foundMatch = strings.Contains(line, pattern)
		}
	}

	if invertProgram {
		return !foundMatch
	} else {
		return foundMatch
	}
}
