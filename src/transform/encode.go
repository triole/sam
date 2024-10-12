package transform

import (
	"encoding/base64"
	"net/url"
)

func (tr Transform) runEncode() (r string) {
	switch tr.Conf.Target {
	case "base64":
		if tr.Conf.Reverse {
			r = tr.fromBase64()
		} else {
			r = tr.toBase64()
		}
	case "url":
		if tr.Conf.Reverse {
			r = tr.fromURL()
		} else {
			r = tr.toURL()
		}
	}
	return
}

func (tr Transform) toBase64() string {
	return base64.StdEncoding.EncodeToString([]byte(tr.Conf.String))
}

func (tr Transform) fromBase64() string {
	r, err := base64.StdEncoding.DecodeString(tr.Conf.String)
	logFatal(err, "can not decode text to base64")
	return string(r)
}

func (tr Transform) toURL() string {
	return url.QueryEscape(tr.Conf.String)
}

func (tr Transform) fromURL() string {
	r, err := url.QueryUnescape(tr.Conf.String)
	logFatal(err, "Error decoding from url to text")
	return r
}
