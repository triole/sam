package transform

import (
	"fmt"
	"strconv"
)

func (tr Transform) runAlign() (r string) {
	if tr.Conf.Target == "right" || tr.Conf.Target == "r" {
		r = tr.alignRight()
	} else {
		r = tr.alignLeft()
	}
	return
}

func (tr Transform) alignLeft() string {
	return fmt.Sprintf("%-"+strconv.Itoa(tr.Conf.Length)+"s", tr.Conf.String)
}

func (tr Transform) alignRight() string {
	return fmt.Sprintf("%"+strconv.Itoa(tr.Conf.Length)+"v", tr.Conf.String)
}
