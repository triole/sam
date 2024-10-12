package transform

import (
	"strings"
)

func (tr Transform) bool() (b bool) {
	b = false
	s := strings.ToLower(tr.Conf.String)
	if s == "true" || s == "enable" || s == "enabled" || s == "1" || s == "on" {
		b = true
	}
	return b
}
