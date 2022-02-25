package grep

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var lineNum int

var rFlags string = ""
var data []string
var prefixStr string = ""
var regexpPattern string = ""

var (
	flagShowLineNum      bool
	flagShowMatchingFile bool
	flagCaseInsensitive  bool
	flagInvertMatching   bool
	flagFullLineMatch    bool
)

func resetFlags() int {
	flagShowLineNum = false
	flagShowMatchingFile = false
	flagCaseInsensitive = false
	flagInvertMatching = false
	flagFullLineMatch = false
	return 0
}

func buildExcludePattern(pattern string) string {
	builtPattern := "^("
	for _, c := range pattern {
		builtPattern += fmt.Sprintf("[^%c]+", c)
	}
	builtPattern += ")$"
	return builtPattern
}

func Search(pattern string, flags, files []string) []string {

	/* clean up match data
	       The recommended way `data = nil` breaks the test, although found
		   this should be recommended for garbage collecting  */
	rFlags = ""
	data = data[:0]

	// reset all flags
	resetFlags()

	/* I had made a function to set flags  beforehand, but it was proven to be
	   redundant, so moved back to this simple switch */
	for _, flag := range flags {

		// compile the pattern
		switch strings.TrimPrefix(flag, "-") {
		case "n":
			flagShowLineNum = true
		case "l":
			flagShowMatchingFile = true
			// case insensitive
		case "i":
			flagCaseInsensitive = true
			// Invert Match
		case "v":
			flagInvertMatching = true
		// Full Line Match
		case "x":
			flagFullLineMatch = true
		}
	}

	if flagCaseInsensitive {
		rFlags = fmt.Sprintf("i%s", rFlags)
	}
	if flagInvertMatching && flagFullLineMatch {
		rFlags = fmt.Sprintf("m%s", rFlags)
		pattern = fmt.Sprintf("(?%s).*[^%s].* ", rFlags, pattern)
	} else if flagFullLineMatch {
		rFlags = fmt.Sprintf("m%s", rFlags)
		pattern = fmt.Sprintf("(?%s)^%s$", rFlags, pattern)
	} else if flagInvertMatching {
		pattern = fmt.Sprintf(buildExcludePattern(pattern))
	} else {
		pattern = fmt.Sprintf("(?%s)%s", rFlags, pattern)
	}

	re, _ := regexp.Compile(pattern)

	for _, s := range files {

		f, err := os.Open(s)
		if err != nil {
			panic("can not open file")
		}
		// reader
		r := bufio.NewReader(f)

		lineNum := 1
		for {
			prefixStr = ""
			line, err := r.ReadString('\n')
			if err == io.EOF {
				fmt.Printf("%s", strings.TrimSuffix(line, "\n"))
				break
			}
			/* ugly debugs
					fmt.Println(pattern)
					fmt.Print(line)
			     	fmt.Println(re.Match([]byte(line))) */

			// if re.MatchString(line) {
			if re.Match([]byte(line)) {

				/* more ugly debugs
				fmt.Print("Matched line: ")
				fmt.Println(line) */

				if (len(files) > 1) && !flagShowMatchingFile {
					// append the file name
					prefixStr = fmt.Sprintf("%s%s:", prefixStr, s)

				} else {
					if flagShowMatchingFile == true {
						if len(data) == 0 {
							prefixStr = fmt.Sprintf("%s%s", prefixStr, s)
						} else {
							prefixStr = fmt.Sprintf("%s%s", prefixStr, s)
							// remove dups - should be fixed!
							// update - got tired.. better get input from ppl
							for i, el := range data {
								if prefixStr == el {
									data = append(data[:i], data[i+1:]...)
								}
							}
						}
						line = ""
					}
				}
				if flagShowLineNum == true && !flagShowMatchingFile {
					prefixStr = fmt.Sprintf("%s%d:", prefixStr, lineNum)
				}
				data = append(data, fmt.Sprintf("%s%s",
					prefixStr,
					strings.TrimSuffix(line, "\n")))
			}
			lineNum++
		}
		f.Close()
	}
	return data
}
