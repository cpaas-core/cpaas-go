package markdown

// implementation to refactor

import (
	"fmt"
	"strings"
)

const (
	headerChar = '#'
	listChar   = '*'
)

var (
	markdownChars = []string{"__", "_"}
	replaceChars  = []string{"<strong>", "</strong>", "<em>", "</em>"}
)

// Render translates markdown to HTML
func Render(markdown string) string {
	// list doesn't need to be a counter
	isList := false
	header := 0
	html := ""

	// replaces the 4 calls to strings.Replace
	for i, replacement := range replaceChars {
		markdown = strings.Replace(markdown, markdownChars[i/2], replacement, 1)
	}

	// add pos and checks to for loop statement to clean up checks and
	// increments
	for pos, prevPos := 0, 0; pos < len(markdown); pos, prevPos = pos+1, pos {
		// use switch case
		switch char, prev := markdown[pos], markdown[prevPos]; char {
		case headerChar:
			header++
		case listChar:
			if !isList {
				html += "<ul>"
				isList = true
			}
			html += "<li>"
		case '\n':
			if isList {
				html += "</li>"
			}
			if header > 0 {
				html += fmt.Sprintf("</h%d>", header)
				header = 0
			}
		// cleans up the extra incrementing code in '*' and '#' cases
		case ' ':
			if prev != listChar && prev != headerChar {
				html += string(char)
			} else if prev == headerChar {
				html += fmt.Sprintf("<h%d>", header)
			}
		default:
			html += string(char)
		}
	}

	if header > 0 {
		return html + fmt.Sprintf("</h%d>", header)
	}
	if isList {
		return html + "</li></ul>"
	}
	return "<p>" + html + "</p>"
}
