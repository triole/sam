package transform

import "strings"

func (tr Transform) Title(str string) string {
	return strings.Title(str)
}
