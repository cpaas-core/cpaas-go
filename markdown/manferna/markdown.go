package markdown

// implementation to refactor

import (
	"fmt"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {

	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)
	header := 0
	list := 0
	html := ""

	for pos, char := range markdown {
		switch char {
		case '#':
			header++
			if markdown[pos+1] != '#' {
				html += fmt.Sprintf("<h%d>", header)
			}
		case '*':
			if list == 0 {
				html += "<ul>"
			}
			html += "<li>"
			list++
		case '\n':
			if list > 0 {
				html += "</li>"
			}
			if header > 0 {
				html += fmt.Sprintf("</h%d>", header)
				header = 0
			}
		default:
			// Delete the whitespace after #
			if char == ' ' && markdown[pos-1] == '#' {
				break
			}
			html += string(char)
		}
	}
	if header > 0 {
		return html + fmt.Sprintf("</h%d>", header)
	}
	if list > 0 {
		return html + "</li></ul>"
	}
	return "<p>" + html + "</p>"
}
