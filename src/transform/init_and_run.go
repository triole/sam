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
	switch tr.Conf.SubCommand {
	// keep-sorted start block=yes
	case "align":
		fmt.Println("Today.")
	case "bool":
		fmt.Println("In two days.")
	case "case":
		fmt.Println("Tomorrow.")
	case "color":
		fmt.Println("In two days.")
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
}
