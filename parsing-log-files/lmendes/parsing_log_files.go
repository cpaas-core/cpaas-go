package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

const logLineRegexp = `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]\s.*`

func IsValidLine(text string) bool {
	re := regexp.MustCompile(logLineRegexp)
	return re.MatchString(text) 
}

func SplitLogLine(text string) []string {
	sectionsRe := regexp.MustCompile(`[a-z]+\s\d?`)
	tagsRe := regexp.MustCompile(`<.*>`)
	parts := sectionsRe.FindAllString(text, -1)
	if len(parts) == 0 {
		parts = tagsRe.FindAllString(text, -1)
		if len(parts) == 0 {
			// I had to "cast" the content to an empty string as when "no match" 
			// the list is empty, not a []string.
			parts = []string{""}
		}
	}
	return parts
}

func CountQuotedPasswords(lines []string) int {
	logLineRegexp := `(?i)\".*password\"`
	matches := 0
	re := regexp.MustCompile(logLineRegexp)
	for _, line := range lines {
		if (re.MatchString(line)) {
			matches++
		}
	}
	return matches
}

func RemoveEndOfLineText(text string) string {
	eolRegexp := regexp.MustCompile(`end-of-line\d+`)
	parts := eolRegexp.Split(text, -1)
	return strings.Join(parts, "")
}

func TagWithUserName(lines []string) []string {
	lineRe := regexp.MustCompile(logLineRegexp)
	userRe := regexp.MustCompile(`User\s+(\w+)`)
	var taggedLines []string
	for _, line := range lines {
		match := lineRe.FindAllString(line, -1)[0]
		userMatch := userRe.FindAllStringSubmatch(match, -1)
		if (len(userMatch) != 0) {
			user := userMatch[0][1]
		    newLine := fmt.Sprintf("[USR] %s %s", user, match)
		    taggedLines = append(taggedLines, newLine)
		} else {
			taggedLines = append(taggedLines, match)
		}
	}
	return taggedLines
}
