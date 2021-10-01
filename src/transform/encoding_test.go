package transform

import "testing"

func TestBase64(t *testing.T) {
	str := "hello world"
	b64 := "aGVsbG8gd29ybGQ="
	assert(tr.ToBase64(str), b64, t)
	assert(tr.FromBase64(b64), str, t)
}
