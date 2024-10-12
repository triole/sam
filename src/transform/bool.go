package transform

import (
	"strings"
)

func (tr Transform) Bool(str string) string {
	s := strings.ToLower(str)
	if s == "true" || s == "enable" || s == "enabled" || s == "1" || s == "on" {
		s = "true"
	} else {
		s = "false"
	}
	return s
}
