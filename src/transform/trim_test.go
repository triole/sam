package transform

import (
	"sam/src/conf"
	"sam/src/implant"
	"testing"
)

func TestTrim(t *testing.T) {
	assertTrim("xhello world", "x", "prefix", false, "hello world", t)
	assertTrim("xxhello world", "xx", "prefix", false, "hello world", t)
	assertTrim("xxhello world", "x", "prefix", false, "xhello world", t)
	assertTrim("xxxhello world", "x", "prefix", true, "hello world", t)
	assertTrim("hello worldx", "x", "suffix", false, "hello world", t)
	assertTrim("hello worldxx", "xx", "suffix", false, "hello world", t)
	assertTrim("hello worldxx", "x", "suffix", true, "hello world", t)
	assertTrim("hello worldxxx", "x", "suffix", true, "hello world", t)
	assertTrim("xhello worldx", "x", "both", false, "hello world", t)
	assertTrim("xxhello worldxx", "x", "both", false, "xhello worldx", t)
	assertTrim("xxxhello worldxxx", "x", "both", true, "hello world", t)
}

func assertTrim(str, sub, target string, agg bool, exp string, t *testing.T) {
	conf := conf.New()
	conf.String = str
	conf.Target = target
	conf.SubString = sub
	conf.Aggressive = agg
	tr := Init(conf, implant.Init())
	assert(conf, tr.runTrim(), exp, t)
}
