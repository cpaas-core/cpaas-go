package parsinglogfiles

import (
	"regexp"
)

//
func IsValidLine(text string) bool {
	// To be considered valid a line should begin with one of the following strings:
	// [TRC], [DBG], [INF], [WRN], [ERR], [FTL]
	re, err := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	if err != nil {
		panic(err)
	}
	b := re.MatchString(text)
	return b
}

func SplitLogLine(text string) []string {
	// any string that has a first character of "<" and a last character of ">" and
	// any combination of the following characters "~", "\*", "=" and "-" in between.
	re, err := regexp.Compile(`<[~*=-]*>`)
	if err != nil {
		panic(err)
	}
	b := re.Split(text, -1)
	return b

}

func CountQuotedPasswords(lines []string) int {
	// Identify log lines where the string "password", which may be in any
	// combination of upper or lower case, is surrounded by quotation marks.
	re, err := regexp.Compile(`(?i)".*password.*"`)
	if err != nil {
		panic(err)
	}
	linesCount := 0
	for _, line := range lines {
		if re.MatchString(line) {
			linesCount = linesCount + 1
		}
	}
	return linesCount
}

func RemoveEndOfLineText(text string) string {
	//take a string and remove the end-of-line text and return a "clean" string
	re, err := regexp.Compile(`end-of-line\d*`)
	if err != nil {
		panic(err)
	}
	return re.ReplaceAllString(text, "")

}

func TagWithUserName(lines []string) []string {
	//- Lines that do not contain the string `"User "` remain unchanged.
	//- For lines that contain the string `"User "`, prefix the line with `[USR]` followed by the user name.
	// You can assume that:
	//User names are followed by at least one whitespace character in the log.
	//There is at most one occurrence of the string `"User "` in each line.
	//User names are non-empty strings that do not contain whitespace.
	re, err := regexp.Compile(`User\s+(\w+)`)
	if err != nil {
		panic(err)
	}
	//outputSlice := make([]string, len(lines))
	outputSlice := []string{}
	for _, line := range lines {
		userMatch := re.FindStringSubmatch(line)
		if userMatch == nil {
			outputSlice = append(outputSlice, line)
		} else {
			lineWithUser := "[USR] " + userMatch[1] + " " + line
			outputSlice = append(outputSlice, lineWithUser)
		}
	}
	return outputSlice

}
