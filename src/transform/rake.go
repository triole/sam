package transform

import (
	"strconv"
	"strings"
)

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
