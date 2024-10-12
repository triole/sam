package transform

import (
	"sam/src/conf"
	"testing"
)

func TestAlign(t *testing.T) {
	conf := conf.InitTest()
	conf.String = "hello"
	conf.Length = 8
	tr = Init(conf)
	assert(tr.alignRight(), "   hello", t)
	assert(tr.alignLeft(), "hello   ", t)
}
