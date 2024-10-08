package transform

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/jzelinskie/whirlpool"
	"lukechampine.com/blake3"
)

func (tr Transform) Md5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (tr Transform) Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (tr Transform) Sha256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (tr Transform) Sha384(str string) string {
	h := sha512.New384()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (tr Transform) Sha512(str string) string {
	h := sha512.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (tr Transform) Blake3(args string) string {
	lenstr, inp := separateFirstArg(args)
	len := parseLengthStr(lenstr)
	str := tr.TrimSpace(inp)
	h := blake3.New(len, nil)
	_, err := h.Write([]byte(str))
	logFatal(err, "Error generating blake3")
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (tr Transform) Whirlpool(str string) string {
	h := whirlpool.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
