package transform

import (
	"fmt"
	"runtime"
	"sam/src/conf"
	"testing"
)

func assert(conf conf.Conf, in, exp string, t *testing.T) {
	if in != exp {
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		callFrom := ""
		if ok && details != nil {
			callFrom = details.Name()
		}
		t.Errorf(
			"\ncall from   %s\ntest input  %s\nassert fail %q != %q",
			callFrom, fmt.Sprintf("%+v", conf), in, exp,
		)
	}
}
