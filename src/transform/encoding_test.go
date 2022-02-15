package transform

import "testing"

func TestBase64(t *testing.T) {
	str := "hello world"
	b64 := "aGVsbG8gd29ybGQ="
	assert(tr.ToBase64(str), b64, t)
	assert(tr.FromBase64(b64), str, t)
}

func TestURL(t *testing.T) {
	str := "hello world, this is good!!"
	url := "hello+world%2C+this+is+good%21%21"
	assert(tr.ToURL(str), url, t)
	assert(tr.FromURL(url), str, t)
}
