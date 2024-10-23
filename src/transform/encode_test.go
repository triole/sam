package transform

import (
	"sam/src/conf"
	"testing"
)

func TestEncode(t *testing.T) {
	assertEncode("hello world", "base64", false, "aGVsbG8gd29ybGQ=", t)
	assertEncode("aGVsbG8gd29ybGQ=", "base64", true, "hello world", t)
	assertEncode("hello world, this is terrific!!=", "url", false, "hello+world%2C+this+is+terrific%21%21%3D", t)
	assertEncode("hello+world%2C+this+is+terrific%21%21%3D", "url", true, "hello world, this is terrific!!=", t)
}

func assertEncode(str, target string, reverse bool, exp string, t *testing.T) {
	conf := conf.New()
	conf.String = str
	conf.Target = target
	conf.Reverse = reverse
	tr := Init(conf)
	assert(conf, tr.runEncode(), exp, t)
}
