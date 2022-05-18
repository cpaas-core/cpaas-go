package markdown

// implementation to refactor

import (
	"fmt"
	"regexp"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	header := 0

	// Using RegExp is more robust
	//markdown = strings.Replace(markdown, "__", "<strong>", 1)
	//markdown = strings.Replace(markdown, "__", "</strong>", 1)
	var stringRE = regexp.MustCompile(`__(.*)__`)
	markdown = stringRE.ReplaceAllString(markdown, "<strong>$1</strong>")

	//markdown = strings.Replace(markdown, "_", "<em>", 1)
	//markdown = strings.Replace(markdown, "_", "</em>", 1)
	var emRE = regexp.MustCompile(`_(.*)_`)
	markdown = emRE.ReplaceAllString(markdown, "<em>$1</em>")

	pos := 0
	list := 0
	html := ""

	// labelled loop for break
charLoop:
	for {
		char := markdown[pos]

		// Remove "pos++" from each case
		pos++

		// using switch instead of a bunch of "if" statements
		switch char {
		case '#':
			for char == '#' {
				header++
				char = markdown[pos]
				pos++
			}
			html += fmt.Sprintf("<h%d>", header)
		case '*':
			if list == 0 {
				html += "<ul>"
			}
			html += "<li>"
			list++
			pos++
		case '\n':
			if list > 0 {
				html += "</li>"
			}
			if header > 0 {
				html += fmt.Sprintf("</h%d>", header)
				header = 0
			}
		default:
			html += string(char)
			if pos >= len(markdown) {
				break charLoop
			}
		}
	}

	// Use switch instead of multiple "if" statements
	// Updated code to use a single return
	switch {
	case header > 0:
		html += fmt.Sprintf("</h%d>", header)
	case list > 0:
		html += "</li></ul>"
	default:
		html = "<p>" + html + "</p>"
	}

	return html
}
