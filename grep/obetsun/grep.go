package grep

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func matchFlag(flagString string, flg string) bool {
	result := false
	flagMatch, errF := regexp.MatchString(flg, flagString)
	if errF == nil && flagMatch {
		result = true
	}
	return result
}

func appendFileName(fileName string, str string) string {
	str = fileName + ":" + str
	return str
}

func appendLineNumber(flagString string, lineNumber int, str string) string {
	if matchFlag(flagString, "-n") {
		str = strconv.Itoa(lineNumber) + ":" + str
	}
	return str
}

func matchStringWithFlags(str string, pattern string, flagStr string, lineNumber int, fileName string) string {
	var result string
	var match bool
	var strLC string
	var patternLC string

	if matchFlag(flagStr, "-i") {
		strLC = strings.ToLower(str)
		patternLC = strings.ToLower(pattern)
	} else {
		strLC = str
		patternLC = pattern
	}

	if matchFlag(flagStr, "-x") {
		match = strLC == patternLC
	} else {
		match = strings.Contains(strLC, patternLC)
	}

	if matchFlag(flagStr, "-v") {
		match = !match
	}

	if match {
		if matchFlag(flagStr, "-l") {
			result = fileName
		} else {
			result = appendLineNumber(flagStr, lineNumber, str)
		}
	}

	return result
}

func arrayContains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func Search(pattern string, flags, files []string) []string {
	result := make([]string, 0)

	for _, f := range files {
		file, err := os.Open(f)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		var flagStr string
		for _, v := range flags {
			flagStr += v
		}

		lineNumber := 0
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			txt := scanner.Text()
			lineNumber++
			res := matchStringWithFlags(txt, pattern, flagStr, lineNumber, f)
			if res != "" {
				if len(files) > 1 && !matchFlag(flagStr, "-l") {
					res = f + ":" + res
				}
				if !arrayContains(result, res) {
					result = append(result, res)
				}
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	return result
}
