package transform

import "testing"

func assert(in, exp string, t *testing.T) {
	if in != exp {
		t.Errorf("assert failed: %q != %q", in, exp)
	}
}
