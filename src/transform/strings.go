package transform

import "strings"

func (tr Transform) Title(str string) string {
	return strings.Title(str)
}

func (tr Transform) Lowercase(str string) string {
	return strings.ToLower(str)
}

func (tr Transform) Uppercase(str string) string {
	return strings.ToUpper(str)
}
