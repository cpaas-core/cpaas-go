// -n Print the line numbers of each matching line.
// -l Print only the names of files that contain at least one matching line.
// -i Match line using a case-insensitive comparison.
// -v Invert the program -- collect all lines that fail to match the pattern.
// -x Only match entire lines, instead of lines that contain a match.
package grep

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Search(pattern string, flags, files []string) []string {

	results := []string{}

	printLineNum := optionEnabled("-n", flags)
	printFilenameOnly := optionEnabled("-l", flags)
	matchCaseInsensetive := optionEnabled("-i", flags)
	matchInverted := optionEnabled("-v", flags)
	matchEntireLine := optionEnabled("-x", flags)

	for _, file := range files {
		fh, err := os.Open(file)
		if err != nil {
			log.Fatalf("failed openning file %s", file)
		}
		lineNum := 1
		scanner := bufio.NewScanner(fh)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Text()
			match := Matcher(matchCaseInsensetive, matchInverted, matchEntireLine, pattern, line)
			if match {
				if printFilenameOnly {
					results = append(results, file)
					break
				}
				if printLineNum {
					line = strconv.Itoa(lineNum) + ":" + line
				}
				if len(files) > 1 {
					line = file + ":" + line
				}
				results = append(results, line)
			}
			lineNum++
		}
		err = fh.Close()
		if err != nil {
			log.Fatalf("failed closing file %s", file)
		}
	}

	return results

}

func optionEnabled(option string, flags []string) bool {
	for _, flag := range flags {
		if option == flag {
			return true
		}
	}
	return false
}

func Matcher(matchCaseInsensetive, matchInverted, matchEntireLine bool, pattern, line string) (match bool) {
	if matchCaseInsensetive {
		pattern = strings.ToLower(pattern)
		line = strings.ToLower(line)
	}
	if matchEntireLine {
		match = pattern == line
	} else {
		match = strings.Contains(line, pattern)
	}
	if matchInverted {
		return !match
	}
	return
}
