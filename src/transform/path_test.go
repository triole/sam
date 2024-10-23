package transform

import (
	"sam/src/conf"
	"testing"
)

func TestPath(t *testing.T) {
	assertPath("/hello//there/world", "dir", "/hello/there", t)
	assertPath("/hello//there/world", "bn", "world", t)
	assertPath("/hello//there/world.zip", "ext", ".zip", t)
	assertPath("/hello//there/world.global", "ext", ".global", t)
}

func assertPath(str, target string, exp string, t *testing.T) {
	conf := conf.New()
	conf.String = str
	conf.Target = target
	tr := Init(conf)
	assert(conf, tr.runPath(), exp, t)
}
