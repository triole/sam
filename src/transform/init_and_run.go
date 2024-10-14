package transform

import (
	"fmt"
	"sam/src/conf"
)

type Transform struct {
	Conf conf.Conf
}

func Init(conf conf.Conf) (tr Transform) {
	tr.Conf = conf
	return
}

func InitTest() (tr Transform) {
	return
}

func (tr Transform) Run() {
	var r string
	switch tr.Conf.SubCommand {
	// keep-sorted start block=yes
	case "align":
		r = tr.runAlign()
	case "bool":
		r = fmt.Sprintf("%v", tr.bool())
	case "calc":
		r = tr.runCalc()
	case "case":
		r = tr.runCase()
	case "color":
		r = tr.runColor()
	case "encode":
		r = tr.runEncode()
	case "hash":
		r = tr.runHash()
	case "path":
		r = tr.runPath()
	case "tidy":
		r = tr.runTidy()
	case "trim":
		r = tr.runTrim()
		// keep-sorted end
	}
	fmt.Printf("%s\n", r)
}
