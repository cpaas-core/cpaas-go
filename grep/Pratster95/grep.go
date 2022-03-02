package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type match struct {
	matches   func(s string) bool
	lines     bool
	filenames bool
}

func parse(pattern string, flags []string) *match {
	var lines bool
	var names bool
	var entire_lines bool
	var caseInsensitive bool
	var inverted bool
	for _, flag := range flags {
		switch flag[1] {
		case 'n':
			lines = true
		case 'l':
			names = true
		case 'i':
			caseInsensitive = true
		case 'v':
			inverted = true
		case 'x':
			entire_lines = true
		}
	}
	if entire_lines {
		pattern = "^" + pattern + "$"
	}
	if caseInsensitive {
		pattern = "(?i)" + pattern
	}
	f := func(s string) bool {
		return regexp.MustCompile(pattern).MatchString(s) != inverted
	}
	return &match{f, lines, names}
}

func ReadLines(file string) []string {
	f, _ := os.Open(file)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	scanner := bufio.NewScanner(f)
	var result []string
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}
func Search(pattern string, flags, files []string) []string {
	match := parse(pattern, flags)
	matches := make([]string, 0)
	for _, file := range files {
		for i, line := range ReadLines(file) {
			if match.matches(line) {
				if match.filenames {
					matches = append(matches, file)
					break
				}
				if match.lines {
					line = fmt.Sprintf("%d:%s", i+1, line)
				}
				if len(files) > 1 {
					line = file + ":" + line
				}
				matches = append(matches, line)
			}
		}
	}
	return matches

}
