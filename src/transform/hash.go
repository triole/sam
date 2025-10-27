package transform

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jameskeane/bcrypt"
	"github.com/jzelinskie/whirlpool"
	"lukechampine.com/blake3"
)

func (tr Transform) runHash() (r string) {
	switch tr.Conf.Target {
	case "md5":
		r = tr.md5()
	case "sha1":
		r = tr.sha1()
	case "sha256":
		r = tr.sha256()
	case "sha384":
		r = tr.sha384()
	case "sha512":
		r = tr.sha512()
	case "blake3":
		r = tr.blake3()
	case "whirlpool":
		r = tr.whirlpool()
	case "rake":
		r = tr.rake()
	case "bcrypt":
		r = tr.bcrypt()
	}
	return
}

func (tr Transform) calculateHash(hasher hash.Hash) {
	if tr.Conf.File != "" {
		fil, err := os.Open(tr.Conf.File)
		if err != nil {
			log.Fatal(err)
		}
		defer fil.Close()
		if _, err := io.Copy(hasher, fil); err != nil {
			log.Fatal(err)
		}
	} else {
		hasher.Write([]byte(tr.Conf.String))
	}
}

func (tr Transform) bcrypt() string {
	salt, _ := bcrypt.Salt(tr.Conf.Rounds)
	hash, _ := bcrypt.Hash(tr.Conf.String, salt)
	return fmt.Sprintf("%s", hash)
}

func (tr Transform) md5() string {
	hasher := md5.New()
	tr.calculateHash(hasher)
	return hex.EncodeToString(hasher.Sum(nil))
}

func (tr Transform) sha1() string {
	hasher := sha1.New()
	tr.calculateHash(hasher)
	bs := hasher.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (tr Transform) sha256() string {
	hasher := sha256.New()
	tr.calculateHash(hasher)
	bs := hasher.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (tr Transform) sha384() string {
	hasher := sha512.New384()
	tr.calculateHash(hasher)
	bs := hasher.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (tr Transform) sha512() string {
	hasher := sha512.New()
	tr.calculateHash(hasher)
	bs := hasher.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (tr Transform) blake3() string {
	hasher := blake3.New(tr.Conf.Length, nil)
	tr.calculateHash(hasher)
	bs := hasher.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (tr Transform) whirlpool() string {
	hasher := whirlpool.New()
	tr.calculateHash(hasher)
	bs := hasher.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (tr Transform) rake() (r string) {
	hash := []rune(tr.blake3())
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
