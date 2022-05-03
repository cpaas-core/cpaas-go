package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	const prefix = "^\\[(TRC|DBG|INF|WRN|ERR|FTL)\\]"
	re := regexp.MustCompile(prefix)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	const split = "<[~*=-]*>"
	re := regexp.MustCompile(split)
	// the int is the number of substrings to return. Negative means return all
	// substrings
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	// (?i) at the beginning of the regexp means case insensitive
	const passwordMatch = "(?i)\".*password.*\""
	re := regexp.MustCompile(passwordMatch)
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	const endOfLine = "end-of-line\\d+"
	re := regexp.MustCompile(endOfLine)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	const usernameMatch = "User[[:space:]]+(\\S+)\\s"
	re := regexp.MustCompile(usernameMatch)
	newLines := make([]string, 0, len(lines))
	for _, line := range lines {
		subMatch := re.FindStringSubmatch(line)
		if subMatch != nil && len(subMatch) > 0 {
			line = "[USR] " + subMatch[1] + " " + line
		}
		newLines = append(newLines, line)
	}
	return newLines
}
