package transform

import (
	"fmt"
	"strings"

	"github.com/gobeam/stringy"
)

func (tr Transform) Title(str string) string {
	return strings.Title(str)
}

func (tr Transform) LowerCase(str string) string {
	return strings.ToLower(str)
}

func (tr Transform) UpperCase(str string) string {
	return strings.ToUpper(str)
}

func (tr Transform) SnakeCase(str string) string {
	sn := stringy.New(str)
	return sn.SnakeCase("?", "").ToLower()
}

func (tr Transform) CamelCase(str string) string {
	sn := stringy.New(str)
	return sn.CamelCase("?", "")
}

func (tr Transform) AlignLeft(args string) string {
	lenstr, inp := separateFirstArg(args)
	return fmt.Sprintf("%-"+lenstr+"v", inp)
}

func (tr Transform) AlignRight(args string) string {
	lenstr, inp := separateFirstArg(args)
	return fmt.Sprintf("%"+lenstr+"v", inp)
}

func (tr Transform) Bool(str string) string {
	s := strings.ToLower(str)
	if s == "true" || s == "enable" || s == "enabled" || s == "1" || s == "on" {
		s = "true"
	} else {
		s = "false"
	}
	return s
}

// TODO: check what to do with this
func wordWrap(text string, lineWidth int) string {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return text
	}
	wrapped := words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}
	return wrapped
}
