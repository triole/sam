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
	case "case":
		r = tr.runCase()
	case "color":
		r = tr.runColor()
	case "encode":
		fmt.Println("In two days.")
	case "hash":
		fmt.Println("In two days.")
	case "path":
		fmt.Println("In two days.")
	case "tidy":
		fmt.Println("In two days.")
	case "trim":
		fmt.Println("In two days.")
	}
	// keep-sorted end
	fmt.Printf("%s\n", r)
}
