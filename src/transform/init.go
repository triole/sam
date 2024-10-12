package transform

import (
	"sam/src/conf"
)

type Transform struct {
	Conf conf.Conf
}

func Init(conf conf.Conf) (tr Transform) {
	tr.Conf = conf
	return
}
