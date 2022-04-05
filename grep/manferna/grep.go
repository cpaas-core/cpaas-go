package grep

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func Search(pattern string, flags, files []string) []string {
	var matchResult []string
	var patternForRegex = pattern
	var invert bool

	// flags that modified the regex -i, -x
	for _, flag := range flags {
		if flag == "-i" {
			patternForRegex = "(?i)" + patternForRegex
			break
		} else if flag == "-x" {
			patternForRegex = "^" + patternForRegex + "$"
			break
			// invert search
		} else if flag == "-v" {
			invert = true
			break
		}
	}
	// for multiples files
	for _, file := range files {
		fileContent, err := os.Open(file) // For read access.
		if err != nil {
			panic(err)
		}

		fileWithMatch := false // for flag "-l"
		lineNumber := 0        // for flag "-n"
		//read line by line
		scanner := bufio.NewScanner(fileContent)

		for scanner.Scan() {
			lineText := scanner.Text()
			lineNumber = lineNumber + 1

			match, _ := regexp.MatchString(patternForRegex, lineText)
			if match {
				if len(flags) == 0 {
					matchResult = append(matchResult, lineText)
				} else {
					for _, flag := range flags {
						if flag == "-n" {
							matchResult = append(matchResult, strconv.Itoa(lineNumber)+":"+lineText)
						} else if flag == "-l" {
							fileWithMatch = true
						} else if flag == "-i" || flag == "-x" {
							matchResult = append(matchResult, lineText)
						}
					}
				}
			} else if !match && invert {
				matchResult = append(matchResult, lineText)
			}
		}
		if fileWithMatch {
			matchResult = append(matchResult, file)
		}
	}
	return matchResult
}
