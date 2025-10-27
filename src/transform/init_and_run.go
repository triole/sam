package transform

import (
	"fmt"
	"sam/src/conf"
	"sam/src/implant"
)

type Transform struct {
	Conf conf.Conf
	Impl implant.Implant
}

func Init(conf conf.Conf, impl implant.Implant) (tr Transform) {
	tr.Conf = conf
	tr.Impl = impl
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
		tr.runColor()
	case "date":
		tr.runDate()
	case "encode":
		r = tr.runEncode()
	case "geo":
		r = tr.runGeo()
	case "hash":
		r = tr.runHash()
	case "match":
		r = tr.runMatch()
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
