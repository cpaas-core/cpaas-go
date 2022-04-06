package railfence

import (
	"strings"
	"unicode"
)

func Encode(message string, rails int) string {
	if message == "" || rails <= 0 {
		return ""
	}
	if rails == 1 {
		return message
	}
	lines := PopulateRails(message, rails)
	out := ""
	for _, line := range lines {
		out += string(line)
	}
	return out
}

func PopulateRails(message string, rails int) [][]rune {
	lines := make([][]rune, rails)
	runes := []rune(message)
	l := len(runes)
	char_cnt := 0
	up := true
	down := false
	line_number := 1

	for char_cnt < l {
		ch := runes[char_cnt]
		if unicode.IsSpace(ch) {
			continue
		}
		lines[line_number-1] = append(lines[line_number-1], ch)
		if line_number == rails {
			up = true
			down = false
		}
		if line_number == 1 {
			down = true
			up = false
		}
		if up {
			line_number--
		}
		if down {
			line_number++
		}
		char_cnt++
	}
	return lines

}

func Decode(message string, rails int) string {
	if message == "" || rails <= 0 {
		return ""
	}
	if rails == 1 {
		return message
	}
	messRunes := []rune(message)
	lines := PopulateRails(strings.Repeat("a", len(messRunes)), rails)
	cnt := 0
	for _, line := range lines {
		for i := range line {
			line[i] = messRunes[cnt]
			cnt++
		}
	}
	out := ""
	char_cnt := 0
	up := true
	down := false
	line_number := 1
	for {
		totallen := 0
		for _, line := range lines {
			totallen += len(line)
		}
		if totallen == 0 {
			break
		}
		l := len(lines[line_number-1])
		if l > 0 {
			out += string(lines[line_number-1][0])
			lines[line_number-1] = lines[line_number-1][1:l]
			char_cnt++
		}
		if line_number == rails {
			up = true
			down = false
		}
		if line_number == 1 {
			down = true
			up = false
		}
		if up {
			line_number--
		}
		if down {
			line_number++
		}
	}
	return out
}
