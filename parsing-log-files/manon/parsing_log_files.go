package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`(?m)^\[(TRC|DBG|INF|WRN|ERR|FTL])]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*-=]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var counter int
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			counter++
		}
	}
	return counter
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	for idx, line := range lines {
		submatch := re.FindStringSubmatch(line)
		if submatch != nil {
			user := submatch[1]
			lines[idx] = fmt.Sprintf("[USR] %s %s", user, line)
		}
	}
	return lines
}
