package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	return regexp.MustCompile("^\\[(INF|TRC|DBG|WRN|ERR|FTL)\\]").MatchString(text)
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile("<[~*=-]*>").Split(text, 3)

}

func CountQuotedPasswords(lines []string) int {
	var reg = regexp.MustCompile(`".*(?i)password.*"`)
	var count int

	for _, v := range lines {
		if reg.MatchString(v) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile("end-of-line[0-9]*").ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var reg = regexp.MustCompile("User\\s+[a-zA-Z0-9]*\\s+")
	var result []string
	for _, v := range lines {
		usr := reg.FindString(v)
		usr = regexp.MustCompile("User\\s+").ReplaceAllString(usr, "[USR] ")
		if usr != "" {
			result = append(result, usr+v)
		} else {

			result = append(result, v)
		}
	}
	return result
}
