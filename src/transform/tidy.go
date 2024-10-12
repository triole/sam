package transform

import "os"

func (tr Transform) runTidy() (r string) {
	switch tr.Conf.Target {
	case "spaces":
		r = tr.spaces()
	case "pseps":
		r = tr.psep()
	}
	return
}

func (tr Transform) psep() string {
	sep := string(os.PathSeparator)
	return rxSub(tr.Conf.String, sep+"+", sep)
}

func (tr Transform) spaces() string {
	return rxSub(tr.Conf.String, `\s+`, " ")
}
