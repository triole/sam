package transform

import (
	"os"
	"regexp"
	"strings"
)

var (
	sep      = string(os.PathSeparator)
	rxScheme = `[^a-zA-Z0-9\-\.` + sep + `]`
)

func (tr Transform) TidyFileName1(str string) (r string) {
	r = tr.sub(str, "["+sep+"]+", "/")
	return
}

func (tr Transform) TidyFileName2(str string) (r string) {
	r = tr.TidyFileName1(str)
	r = tr.sub(r, rxScheme, "_")
	return
}

func (tr Transform) TidyFileName3(str string) (r string) {
	r = tr.TidyFileName2(str)
	r = strings.ToLower(r)
	return
}

func (tr Transform) TidyFileName4(str string) (r string) {
	r = tr.TidyFileName3(str)
	r = tr.removeMultiples(r)
	return
}

func (tr Transform) removeMultiples(s string) (r string) {
	r = tr.sub(s, "[_]+", "_")
	r = tr.sub(r, "[-]+", "-")
	return
}

func (tr Transform) sub(str string, rx string, rep string) (r string) {
	re := regexp.MustCompile(rx)
	r = re.ReplaceAllString(str, rep)
	return
}
