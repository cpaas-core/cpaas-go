package markdown

// implementation to refactor

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	headerRegexp    = regexp.MustCompile(`^(#+)[[:blank:]](.*)`)
	listItemRegexp  = regexp.MustCompile(`(?m:^\*[[:blank:]].*$)`)
	listRegexp      = regexp.MustCompile(`(?s:<li>.*</li>)`)
	paragraphRegexp = regexp.MustCompile(`(\n?)([^\n]+)(\n?)`)
)

// Render translates markdown to HTML
func Render(markdown string) string {
	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)

	markdown = headerRegexp.ReplaceAllStringFunc(markdown, func(s string) string {
		headerLevel := strings.Count(s, "#")
		s = strings.Trim(strings.Replace(s, "#", "", -1), " ")
		return fmt.Sprintf("<h%d>%s</h%d>", headerLevel, s, headerLevel)
	})

	markdown = listItemRegexp.ReplaceAllStringFunc(markdown, func(s string) string {
		s = strings.Replace(s, "* ", "", -1)
		return fmt.Sprintf("<li>%s</li>", s)
	})

	markdown = listRegexp.ReplaceAllStringFunc(markdown, func(s string) string {
		return fmt.Sprintf("<ul>%s</ul>", s)
	})

	markdown = paragraphRegexp.ReplaceAllStringFunc(markdown, func(s string) string {
		if strings.HasPrefix(s, "<h") || strings.HasPrefix(s, "<ul") || strings.HasPrefix(s, "<li") {
			return s
		}
		return fmt.Sprintf("<p>%s</p>", s)
	})

	markdown = strings.ReplaceAll(markdown, "\n", "")
	return markdown
}
