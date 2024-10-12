package transform

func (tr Transform) ReplaceMultiSpace(str string) string {
	return tr.sub(str, `\s+`, " ")
}
