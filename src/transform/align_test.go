package transform

import (
	"sam/src/conf"
	"testing"
)

func TestAlign(t *testing.T) {
	assertAlign("hello", "right", "   hello", t)
	assertAlign("hello", "r", "   hello", t)
	assertAlign("hello", "left", "hello   ", t)
	assertAlign("hello", "l", "hello   ", t)
}

func assertAlign(str, target, exp string, t *testing.T) {
	conf := conf.New()
	conf.String = str
	conf.Target = target
	conf.Length = 8
	tr := Init(conf)
	assert(conf, tr.runAlign(), exp, t)
}
