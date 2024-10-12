package transform

import (
	"testing"
)

func TestAlign(t *testing.T) {
	assert(tr.AlignRight("8 hello"), "   hello", t)
	assert(tr.AlignLeft("8 hello"), "hello   ", t)
}
