package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) (numLines int) {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			numLines++
		}
	}
	return
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) (updatedLines []string) {
	re := regexp.MustCompile(`User\s+(\S+)`)
	for _, line := range lines {
		match := re.FindAllStringSubmatch(line, 1) // slice of slices
		if len(match) > 0 {
			updatedLines = append(updatedLines, fmt.Sprintf("[USR] %s %s", match[0][1], line))
		} else {
			updatedLines = append(updatedLines, line)
		}
	}
	return
}
