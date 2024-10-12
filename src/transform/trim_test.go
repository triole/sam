package transform

import "testing"

func TestTrim(t *testing.T) {
	// assert(tr.TrimPrefix("x xhello world"), "hello world", t)
	// assert(tr.TrimPrefix("xx xxhello world"), "hello world", t)
	// assert(tr.TrimPrefix("x xxhello world"), "xhello world", t)
	// assert(tr.TrimPrefixAggressive("x xxxhello world"), "hello world", t)

	// assert(tr.TrimSuffix("x hello worldx"), "hello world", t)
	// assert(tr.TrimSuffix("xx hello worldxx"), "hello world", t)
	// assert(tr.TrimSuffix("x hello worldxx"), "hello worldx", t)
	// assert(tr.TrimSuffixAggressive("x hello worldxxx"), "hello world", t)

	// assert(tr.TrimSpace(" hello world "), "hello world", t)
	// assert(tr.TrimSpace(" hello world  "), "hello world", t)
	// assert(tr.TrimSpace("\t   hello world\t"), "hello world", t)
}
