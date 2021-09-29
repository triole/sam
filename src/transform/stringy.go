package transform

import (
	"strconv"
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

func (tr Transform) Bool(str string) string {
	s := strings.ToLower(str)
	if s == "true" || s == "enable" || s == "enabled" || s == "1" || s == "on" {
		s = "on"
	} else {
		s = "off"
	}
	sn := stringy.New(s)
	return strconv.FormatBool(sn.Boolean())
}
