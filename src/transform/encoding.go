package transform

import (
	"encoding/base64"
	"net/url"
)

func (tr Transform) ToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func (tr Transform) FromBase64(str string) string {
	r, err := base64.StdEncoding.DecodeString(str)
	logFatal(err, "Error decoding from text to base64")
	return string(r)
}

func (tr Transform) ToURL(str string) string {
	return url.QueryEscape(str)
}

func (tr Transform) FromURL(str string) string {
	r, err := url.QueryUnescape(str)
	logFatal(err, "Error decoding from url to text")
	return r
}
