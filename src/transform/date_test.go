package transform

import (
	"sam/src/conf"
	"sam/src/implant"
	"strconv"
	"strings"
	"testing"
	"time"
)

type timeZoneStrings struct {
	Abbreviation string
	Offset       string
	OffsetColon  string
}

func TestDate(t *testing.T) {
	// map key is the expected result, list contains test input values
	inp := make(map[string][]string)
	inp["Wed Jan 22 15:33:00 [TZ_ABB] 2025"] = []string{
		"Wed Jan 22 15:33:00 [TZ_ABB] 2025",
		"2025-01-22T15:33:00[TZ_COL]",
		"2025-01-22T15:33:00.00000000[TZ_COL]",
		"22 Jan 25 15:33 [TZ_OFF]",
		"Wednesday, 22-Jan-25 15:33:00 [TZ_ABB]",
		"Wed, 22 Jan 2025 15:33:00 [TZ_OFF]",
		"2025-01-22T15:33:00",
		"2025-01-22t15:33:00",
	}
	inp["Sun Jun 22 15:33:00 [TZ_ABB] 2025"] = []string{
		"Sun Jun 22 15:33:00 [TZ_ABB] 2025",
		"2025-06-22T15:33:00[TZ_COL]",
		"2025-06-22T15:33:00.00000000[TZ_COL]",
		"22 Jun 25 15:33 [TZ_OFF]",
		"Sunday, 22-Jun-25 15:33:00 [TZ_ABB]",
		"Sun, 22 Jun 2025 15:33:00 [TZ_OFF]",
		"2025-06-22T15:33:00",
		"2025-06-22t15:33:00",
	}
	for expectedString, el := range inp {
		inpList := el[1:]
		for _, input := range inpList {
			exString := replaceTimeZonesStrings(expectedString)
			inString := replaceTimeZonesStrings(input)
			assertDate(inString, exString, t)
		}
	}
}

func assertDate(str, expString string, t *testing.T) {
	conf := conf.New()
	conf.String = str
	tr := Init(conf, implant.Init())
	dat, _ := tr.strToDate()
	exp, _ := time.ParseInLocation(time.UnixDate, expString, time.Local)
	assert(
		conf,
		strconv.Itoa(int(dat.Unix())),
		strconv.Itoa(int(exp.Unix())),
		t,
	)
}

func getCorrectTimeZones(month time.Month) (tzs timeZoneStrings) {
	tim := time.Date(2025, month, 22, 15, 33, 00, 00, time.Local)
	tzs.Abbreviation = tim.Format("MST")
	tzs.Offset = tim.Format("-0700")
	tzs.OffsetColon = tim.Format("-07:00")
	return
}

func replaceTimeZonesStrings(str string) (r string) {
	winterTZ := getCorrectTimeZones(1)
	summerTZ := getCorrectTimeZones(6)
	r = strings.Replace(str, "[TZ_ABB]", summerTZ.Abbreviation, -1)
	r = strings.Replace(r, "[TZ_OFF]", summerTZ.Offset, -1)
	r = strings.Replace(r, "[TZ_COL]", summerTZ.OffsetColon, -1)
	if strings.Contains(str, "Jan") || strings.Contains(str, "2025-01-") {
		r = strings.Replace(str, "[TZ_ABB]", winterTZ.Abbreviation, -1)
		r = strings.Replace(r, "[TZ_OFF]", winterTZ.Offset, -1)
		r = strings.Replace(r, "[TZ_COL]", winterTZ.OffsetColon, -1)
	}
	return
}
