package transform

import "strings"

func separateFirstArg(s string) (string, string) {
	arr := strings.Split(s, " ")
	return arr[0], strings.Join(arr[1:], " ")
}
