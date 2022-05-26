package markdown

// implementation to refactor

import (
	"fmt"
	"regexp"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {

	subs := regexp.MustCompile(`(.*)__(.*)__(.*)`)
	markdown = subs.ReplaceAllStringFunc(markdown, func(s string) string {
		res := subs.FindStringSubmatch(s)
		return fmt.Sprintf("%s<strong>%s</strong>%s", res[1], res[2], res[3])
	})

	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)

	if regexp.MustCompile(`[\*#]`).FindString(markdown) == "" {
		return "<p>" + markdown + "</p>"

	}

	markdown = regexp.MustCompile(`^(#+)\s(.*)\n*`).ReplaceAllStringFunc(markdown, func(s string) string {
		hCount := strings.Count(s, "#")
		s = strings.TrimSpace(strings.Replace(s, "#", "", -1))
		return fmt.Sprintf("<h%d>%s</h%d>", hCount, s, hCount)
	})

	regList := regexp.MustCompile(`(\*\s+([^\*]+))`)

	markdown = regList.ReplaceAllStringFunc(markdown, func(s string) string {
		return fmt.Sprintf("<li>%s</li>", strings.TrimSpace(strings.Replace(s, "* ", "", -1)))
	})

	regUl := regexp.MustCompile(`(<li>.*</li>)`)

	markdown = regUl.ReplaceAllStringFunc(markdown, func(s string) string {
		return fmt.Sprintf("<ul>%s</ul>", s)
	})

	return markdown
}
