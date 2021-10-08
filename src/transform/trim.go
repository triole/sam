package transform

import "strings"

func (tr Transform) TrimSuffix(args string) string {
	suf, str := separateFirstArg(args)
	return strings.TrimSuffix(str, suf)
}

func (tr Transform) TrimSuffixAggressive(args string) string {
	suf, str := separateFirstArg(args)
	for strings.HasSuffix(str, suf) == true {
		str = strings.TrimSuffix(str, suf)
	}
	return str
}

func (tr Transform) TrimPrefixAggressive(args string) string {
	pre, str := separateFirstArg(args)
	for strings.HasPrefix(str, pre) == true {
		str = strings.TrimPrefix(str, pre)
	}
	return str
}

func (tr Transform) TrimPrefix(args string) string {
	pre, str := separateFirstArg(args)
	return strings.TrimPrefix(str, pre)
}

func (tr Transform) TrimSpace(str string) string {
	return strings.TrimSpace(str)
}
func (tr Transform) RemoveMultiSpace(str string) string {
	return tr.sub(str, `\s+`, " ")
}
