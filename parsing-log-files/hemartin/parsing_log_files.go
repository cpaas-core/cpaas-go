package parsinglogfiles

import (
	"fmt"
	"regexp"
)

var (
	validLineRegexp = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`)
	separatorRegexp = regexp.MustCompile(`<([.\*~=-]*)>`)
	passwordRegexp  = regexp.MustCompile(`".*(?i:password)+.*"`)
	EOLRegexp       = regexp.MustCompile(`(end-of-line[0-9]+)`)
	UserRegexp      = regexp.MustCompile(`User\s+(\w*)\s.*`)
)

func IsValidLine(text string) bool {
	return validLineRegexp.MatchString(text)
}

func SplitLogLine(text string) []string {
	return separatorRegexp.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	nPasswords := 0
	for _, line := range lines {
		if passwordRegexp.MatchString(line) {
			nPasswords += 1
		}
	}
	return nPasswords
}

func RemoveEndOfLineText(text string) string {
	if EOLRegexp.MatchString(text) {
		return EOLRegexp.ReplaceAllString(text, "")
	}
	return text
}

func TagWithUserName(lines []string) []string {
	outputLines := make([]string, len(lines))

	for i, line := range lines {
		outputLines[i] = line
		if UserRegexp.MatchString(line) {
			user := UserRegexp.FindAllStringSubmatch(line, -1)[0][1]
			outputLines[i] = fmt.Sprintf("[USR] %s %s", user, line)
		}
	}

	return outputLines
}
