package transform

import (
	"testing"
)

func TestSnakeCase(t *testing.T) {
	assert(tr.SnakeCase("hello world"), "hello_world", t)
	assert(tr.SnakeCase("helloWorld"), "hello_world", t)
	assert(tr.SnakeCase("hello   WORLD"), "hello_world", t)

	assert(tr.CamelCase("hello_world"), "HelloWorld", t)
	assert(tr.CamelCase("helloWorld"), "HelloWorld", t)
	assert(tr.CamelCase("hello   WORLD"), "HelloWORLD", t)

	assert(tr.Bool("true"), "true", t)
	assert(tr.Bool("false"), "false", t)
	assert(tr.Bool("enable"), "true", t)
	assert(tr.Bool("disable"), "false", t)
	assert(tr.Bool("enabled"), "true", t)
	assert(tr.Bool("disabled"), "false", t)
	assert(tr.Bool("on"), "true", t)
	assert(tr.Bool("off"), "false", t)
	assert(tr.Bool("1"), "true", t)
	assert(tr.Bool("0"), "false", t)
	assert(tr.Bool("any_other_string"), "false", t)
}
