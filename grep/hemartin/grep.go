package grep

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
)

const printLinesFlag = "-n"
const printNamesFlag = "-l"
const caseInsensitiveFlag = "-i"
const invertFlag = "-v"
const matchLinesFlag = "-x"

func Search(pattern string, flags, files []string) []string {
	result := []string{}

	printLines := sliceContains(flags, printLinesFlag)
	caseInsensitive := sliceContains(flags, caseInsensitiveFlag)
	printNames := sliceContains(flags, printNamesFlag)
	matchLines := sliceContains(flags, matchLinesFlag)
	invert := sliceContains(flags, invertFlag)

	regexFlags := ""
	if caseInsensitive {
		regexFlags += "(?i)"
	}

	if matchLines {
		pattern = fmt.Sprintf("^%s$", pattern)
	}

	regexPattern := fmt.Sprintf(`%s%s`, regexFlags, pattern)

	compiledPattern := regexp.MustCompile(regexPattern)
	filesContents := readFiles(files)
	for _, fileName := range files {
		lines := filesContents[fileName]
		for lineIndex, line := range lines {
			if string(line) == "" {
				continue
			}

			match := compiledPattern.Match(line)
			if (match && !invert) || (!match && invert) {

				prefix := ""
				if len(files) > 1 {
					prefix += fmt.Sprintf("%s:", fileName)
				}

				if printLines {
					prefix += fmt.Sprintf("%d:", lineIndex+1)
				}

				matchResult := fmt.Sprintf("%s%s", prefix, line)
				if printNames {
					if !sliceContains(result, fileName) {
						result = append(result, fileName)
					}
				} else {
					result = append(result, matchResult)
				}
			}
		}
	}

	return result
}

func sliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func readFiles(files []string) map[string][][]byte {
	m := make(map[string][][]byte)

	for _, file := range files {
		text, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}

		m[file] = bytes.Split(text, []byte("\n"))
	}

	return m
}
