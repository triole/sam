package transform

import (
	"encoding/base64"
	"log"
)

func (tr Transform) ToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func (tr Transform) FromBase64(str string) string {
	r, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatalf("Error decoding base64 string: %s", err.Error())
	}
	return string(r)
}
