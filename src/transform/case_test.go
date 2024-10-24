package transform

import (
	"sam/src/conf"
	"sam/src/implant"
	"testing"
)

func TestCase(t *testing.T) {
	assertCase("heLLo_woRLD", "lower", "hello_world", t)
	assertCase("heLLo_woRLD", "upper", "HELLO_WORLD", t)
	assertCase("hello_world", "camel", "HelloWorld", t)
	assertCase("helloWorld", "camel", "HelloWorld", t)
	assertCase("hello   WORLD", "camel", "HelloWORLD", t)
	assertCase("hello world", "snake", "hello_world", t)
	assertCase("helloWorld", "snake", "hello_world", t)
	assertCase("hello   WORLD", "snake", "hello_world", t)
	assertCase("hello my friend, this is world", "title", "Hello My Friend, This Is World", t)
}

func assertCase(str, target, exp string, t *testing.T) {
	conf := conf.New()
	conf.String = str
	conf.Target = target
	conf.Language = "english"
	tr := Init(conf, implant.Init())
	assert(conf, tr.runCase(), exp, t)
}
