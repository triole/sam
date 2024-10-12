package transform

import "strings"

func (tr Transform) runTrim() (r string) {
	switch tr.Conf.Target {
	case "prefix":
		r = tr.trimPrefix()
	case "suffix":
		r = tr.trimSuffix()
	case "both":
		r = tr.trimBoth()
	}
	return
}

func (tr Transform) trimPrefix() string {
	if tr.Conf.Aggressive {
		return tr.trimPrefixAggressive()
	}
	return strings.TrimPrefix(tr.Conf.String, tr.Conf.SubString)
}

func (tr Transform) trimPrefixAggressive() string {
	for strings.HasPrefix(tr.Conf.String, tr.Conf.SubString) {
		tr.Conf.String = strings.TrimPrefix(tr.Conf.String, tr.Conf.SubString)
	}
	return tr.Conf.String
}

func (tr Transform) trimSuffix() string {
	if tr.Conf.Aggressive {
		return tr.trimSuffixAggressive()
	}
	return strings.TrimSuffix(tr.Conf.String, tr.Conf.SubString)
}

func (tr Transform) trimSuffixAggressive() string {
	for strings.HasSuffix(tr.Conf.String, tr.Conf.SubString) {
		tr.Conf.String = strings.TrimSuffix(tr.Conf.String, tr.Conf.SubString)
	}
	return tr.Conf.String
}

func (tr Transform) trimBoth() string {
	tr.Conf.String = tr.trimPrefix()
	return tr.trimSuffix()
}
