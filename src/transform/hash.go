package transform

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

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

// *** rake part below ***
func (tr Transform) Rake(args string) (r string) {
	hash := []rune(tr.Blake3(args))
	for pos, char := range hash {
		r += tr.toNewChar(pos, int(char))
	}
	return r
}

func (tr Transform) toNewChar(i, c int) (r string) {
	shift := i*c + c*c
	for i := 4; i <= 100; i = i + 4 {
		if shift%i == 0 {
			r = strconv.Itoa(tr.crossSum(i))
		}
	}
	if r == "" {
		r = tr.intToLetters(shift)
	}
	return
}

func (tr Transform) intToLetters(number int) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += tr.intToLetters(firstLetter)
		letters += string(rune('A' + number%26))
	} else {
		letters += string(rune('A' + number))
	}
	if number%2 == 0 {
		letters = strings.ToLower(letters)
	}
	return
}

func (tr Transform) crossSum(i int) (r int) {
	r = tr.sumDigits(i)
	if r > 9 {
		r = tr.crossSum(r)
	}
	return
}

func (tr Transform) sumDigits(number int) int {
	remainder := 0
	sumResult := 0
	for number != 0 {
		remainder = number % 10
		sumResult += remainder
		number = number / 10
	}
	return sumResult
}
