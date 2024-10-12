package transform

func (tr Transform) ReplaceMultiSpace(str string) string {
	return rxSub(str, `\s+`, " ")
}
