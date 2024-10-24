package transform

import (
	"sam/src/conf"
	"sam/src/implant"
	"testing"
)

func TestTidy(t *testing.T) {
	assertTidy("   \thello world\t", "spaces", " hello world ", t)
	assertTidy("/hello///world/", "pseps", "/hello/world/", t)
}

func assertTidy(str, target, exp string, t *testing.T) {
	conf := conf.New()
	conf.String = str
	conf.Target = target
	tr := Init(conf, implant.Init())
	assert(conf, tr.runTidy(), exp, t)
}
