package transform

import (
	"sam/src/conf"
	"strconv"
	"testing"
)

func TestDate(t *testing.T) {
	exp := 1737552791
	assertDate("Wed Jan 22 15:33:11 CEST 2025", exp, t)
	assertDate("2025-01-22T15:33:11+02:00", exp, t)
	assertDate("2025-01-22T15:33:11.00000000+02:00", exp, t)
	assertDate("22 Jan 25 15:33 +0200", exp-11, t)
	assertDate("Wednesday, 22-Jan-25 15:33:11 CEST", exp, t)
	assertDate("Wed, 22 Jan 2025 15:33:11 +0200", exp, t)
	// assertDate("2025-01-22T15:33:11", exp, t)
}

func assertDate(str string, exp int, t *testing.T) {
	conf := conf.New()
	conf.String = str
	tr := Init(conf)
	tr.loadLayouts()
	dat, _ := tr.strToDate()
	assert(conf, strconv.Itoa(int(dat.Unix())), strconv.Itoa(exp), t)
}
