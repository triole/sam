package transform

import "strings"

func (tr Transform) title(str string) string {
	return strings.Title(str)
}
