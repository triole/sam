package transform

import (
	"sam/src/conf"
	"sam/src/implant"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	winterTimeExp := prepExpectationTime("Wed Jan 22 15:33:11 CET 2025")
	assertDate("Wed Jan 22 15:33:11 CET 2025", winterTimeExp, t)
	assertDate("2025-01-22T15:33:11+01:00", winterTimeExp, t)
	assertDate("2025-01-22T15:33:11.00000000+01:00", winterTimeExp, t)
	assertDate("22 Jan 25 15:33 +0100", winterTimeExp-11, t)
	assertDate("Wednesday, 22-Jan-25 15:33:11 CET", winterTimeExp, t)
	assertDate("Wed, 22 Jan 2025 15:33:11 +0100", winterTimeExp, t)
	assertDate("2025-01-22T15:33:11", winterTimeExp, t)
	assertDate("2025-01-22t15:33:11", winterTimeExp, t)

	summerTimeExp := prepExpectationTime("Sun Jun 22 15:33:11 CEST 2025")
	assertDate("Sun Jun 22 15:33:11 CEST 2025", summerTimeExp, t)
	assertDate("2025-06-22T15:33:11+02:00", summerTimeExp, t)
	assertDate("2025-06-22T15:33:11.00000000+02:00", summerTimeExp, t)
	assertDate("22 Jun 25 15:33 +0200", summerTimeExp-11, t)
	assertDate("Sunday, 22-Jun-25 15:33:11 CEST", summerTimeExp, t)
	assertDate("Sun, 22 Jun 2025 15:33:11 +0200", summerTimeExp, t)
	assertDate("2025-06-22T15:33:11", summerTimeExp, t)
	assertDate("2025-06-22t15:33:11", summerTimeExp, t)
}

func assertDate(str string, exp int, t *testing.T) {
	conf := conf.New()
	conf.String = str
	tr := Init(conf, implant.Init())
	dat, _ := tr.strToDate()
	assert(conf, strconv.Itoa(int(dat.Unix())), strconv.Itoa(exp), t)
}

func prepExpectationTime(str string) (ts int) {
	tim, _ := time.ParseInLocation(
		time.UnixDate, strings.ToUpper(str), time.Local,
	)
	ts = int(tim.Unix())
	return
}
