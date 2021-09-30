package transform

import "testing"

var (
	tr = Init()
)

func assert(in, exp string, t *testing.T) {
	if in != exp {
		t.Errorf("Assert failed: %q != %q", in, exp)
	}
}
