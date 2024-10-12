package transform

import (
	"fmt"
)

func (tr Transform) AlignLeft(args string) string {
	lenstr, inp := separateFirstArg(args)
	return fmt.Sprintf("%-"+lenstr+"v", inp)
}

func (tr Transform) AlignRight(args string) string {
	lenstr, inp := separateFirstArg(args)
	return fmt.Sprintf("%"+lenstr+"v", inp)
}
