package transform

import (
	"os"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var (
	sep      = string(os.PathSeparator)
	rxScheme = `[^a-zA-Z0-9\-\.` + sep + `]`
)

func (tr Transform) DirName(str string) (r string) {
	r = tr.trimSuffixAggressive(str, sep)
	r = tr.find("^.*"+sep, r)
	r = strings.TrimSuffix(r, sep)
	return
}

func (tr Transform) TidyFileName1(str string) (r string) {
	r = tr.sub(str, "["+sep+"]+", "/")
	return
}

func (tr Transform) TidyFileName2(str string) (r string) {
	r = tr.TidyFileName1(str)
	r = tr.specialCharacterTreatment(r)
	r = tr.removeAccents(r)
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

func (tr Transform) specialCharacterTreatment(s string) (r string) {
	r = s
	r = strings.Replace(r, "ä", "ae", -1)
	r = strings.Replace(r, "Ä", "Ae", -1)
	r = strings.Replace(r, "ö", "oe", -1)
	r = strings.Replace(r, "Ö", "Oe", -1)
	r = strings.Replace(r, "ü", "ue", -1)
	r = strings.Replace(r, "Ü", "Ue", -1)
	r = strings.Replace(r, "ß", "ss", -1)
	return
}

func (tr Transform) removeAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, s)
	if e != nil {
		panic(e)
	}
	return output
}

func (tr Transform) removeMultiples(s string) (r string) {
	r = tr.sub(s, "[_]+", "_")
	r = tr.sub(r, "[-]+", "-")
	return
}

func (tr Transform) trimSuffixAggressive(str, suf string) string {
	for strings.HasSuffix(str, suf) == true {
		str = strings.TrimSuffix(str, suf)
	}
	return str
}

// regex functions
// TODO: maybe move into an own package later
func (tr Transform) compile(str string) (r *regexp.Regexp) {
	r, _ = regexp.Compile(str)
	return
}

func (tr Transform) find(rx string, content string) (r string) {
	temp := tr.compile(rx)
	r = temp.FindString(content)
	return
}

func (tr Transform) sub(str string, rx string, rep string) (r string) {
	re := regexp.MustCompile(rx)
	r = re.ReplaceAllString(str, rep)
	return
}
