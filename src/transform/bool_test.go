package transform

import (
	"fmt"
	"sam/src/conf"
	"sam/src/implant"
	"testing"
)

func TestBool(t *testing.T) {
	assertBool("true", true, t)
	assertBool("false", false, t)
	assertBool("enable", true, t)
	assertBool("disable", false, t)
	assertBool("enabled", true, t)
	assertBool("disabled", false, t)
	assertBool("on", true, t)
	assertBool("off", false, t)
	assertBool("1", true, t)
	assertBool("0", false, t)
	assertBool("any_other_string", false, t)
}

func assertBool(str string, exp bool, t *testing.T) {
	conf := conf.New()
	conf.String = str
	tr := Init(conf, implant.Init())
	res := fmt.Sprintf("%v", tr.bool())
	assert(conf, res, fmt.Sprintf("%v", exp), t)
}
