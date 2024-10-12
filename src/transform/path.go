package transform

import (
	"path"
)

func (tr Transform) runPath() (r string) {
	switch tr.Conf.Target {
	case "dir":
		r = tr.dir()
	case "bn":
		r = tr.basename()
	case "ext":
		r = tr.ext()
	}
	return
}

func (tr Transform) dir() (r string) {
	return path.Dir(tr.Conf.String)
}

func (tr Transform) basename() (r string) {
	return path.Base(tr.Conf.String)
}

func (tr Transform) ext() (r string) {
	return path.Ext(tr.Conf.String)
}
