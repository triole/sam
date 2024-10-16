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
		r = tr.lowerCase(lang)
	case "upper":
		r = tr.upperCase(lang)
	case "camel":
		r = tr.camelCase()
	case "snake":
		r = tr.snakeCase()
	case "title":
		r = tr.titleCase(lang)
	}
	return
}

func (tr Transform) lowerCase(lang language.Tag) string {
	return cases.Lower(lang).String(tr.Conf.String)
}

func (tr Transform) upperCase(lang language.Tag) string {
	return cases.Upper(lang).String(tr.Conf.String)
}

func (tr Transform) camelCase() string {
	sn := stringy.New(tr.Conf.String)
	return sn.CamelCase("?", "")
}

func (tr Transform) snakeCase() string {
	sn := stringy.New(tr.Conf.String)
	return sn.SnakeCase("?", "").ToLower()
}

func (tr Transform) titleCase(lang language.Tag) string {
	return cases.Title(lang).String(tr.Conf.String)
}
