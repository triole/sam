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
}
