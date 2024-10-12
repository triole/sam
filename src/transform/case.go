package transform

import (
	"github.com/gobeam/stringy"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (tr Transform) runCase() (r string) {
	var lang language.Tag
	if tr.Conf.SubCommand == "lower" ||
		tr.Conf.SubCommand == "upper" ||
		tr.Conf.SubCommand == "title" {
		lang = []language.Tag{
			language.MustParse(tr.Conf.Language),
		}[0]
	}
	switch tr.Conf.Target {
	case "lower":
		r = tr.LowerCase(lang)
	case "upper":
		r = tr.UpperCase(lang)
	case "camel":
		r = tr.CamelCase()
	case "snake":
		r = tr.SnakeCase()
	case "title":
		r = tr.TitleCase(lang)
	}
	return
}

func (tr Transform) LowerCase(lang language.Tag) string {
	return cases.Lower(lang).String(tr.Conf.String)
}

func (tr Transform) UpperCase(lang language.Tag) string {
	return cases.Upper(lang).String(tr.Conf.String)
}

func (tr Transform) CamelCase() string {
	sn := stringy.New(tr.Conf.String)
	return sn.CamelCase("?", "")
}

func (tr Transform) SnakeCase() string {
	sn := stringy.New(tr.Conf.String)
	return sn.SnakeCase("?", "").ToLower()
}

func (tr Transform) TitleCase(lang language.Tag) string {
	return cases.Title(lang).String(tr.Conf.String)
}
