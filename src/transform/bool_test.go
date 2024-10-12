package transform

import (
	"testing"
)

func TestBool(t *testing.T) {
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
