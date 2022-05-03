package parsinglogfiles

import (
	"regexp"
)

var validLineRegExp = regexp.MustCompile("^\\[(TRC|DBG|INF|WRN|ERR|FTL)\\]")

func IsValidLine(text string) bool {
	return validLineRegExp.MatchString(text)
}

var splitLogLineRegExp = regexp.MustCompile("<(~|\\*|=|-)*>")

func SplitLogLine(text string) []string {
	return splitLogLineRegExp.Split(text, -1)
}

var quotedStringRegEx = regexp.MustCompile("\".*\"")
var passwordRegExp = regexp.MustCompile("(?i)password")

func CountQuotedPasswords(lines []string) int {
	numberMatchingLines := 0
	for _, line := range lines {
		inQuotes := quotedStringRegEx.FindString(line)
		if passwordRegExp.MatchString(inQuotes) {
			numberMatchingLines++
		}
	}
	return numberMatchingLines
}

var endOfLineRegEx = regexp.MustCompile("end-of-line(\\d)*")

func RemoveEndOfLineText(text string) string {
	return endOfLineRegEx.ReplaceAllString(text, "")
}

var userNameRegEx = regexp.MustCompile("User +([[:alnum:]]+)")

func TagWithUserName(lines []string) []string {
	newLines := []string{}
	for _, line := range lines {
		if userNameRegEx.MatchString(line) {
			userName := userNameRegEx.FindStringSubmatch(line)
			newLine := "[USR] " + userName[1] + " " + line
			newLines = append(newLines, newLine)
		} else {
			newLines = append(newLines, line)
		}
	}
	return newLines
}
