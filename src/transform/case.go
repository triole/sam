package transform

import (
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
