package transform

import "testing"

func TestTidy(t *testing.T) {
	assert(tr.ReplaceMultiSpace("   \thello world\t"), " hello world ", t)
}
