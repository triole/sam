package transform

import "testing"

func TestParseLenghtStr(t *testing.T) {
	validateTestParseLenghtStr("64", 64, t)
	validateTestParseLenghtStr("1000", 1000, t)
	validateTestParseLenghtStr("1e3", 1000, t)
	validateTestParseLenghtStr("4p4", 256, t)
}

func validateTestParseLenghtStr(inp string, exp int, t *testing.T) {
	res := parseLengthStr(inp)
	if res != exp {
		t.Errorf(
			"testParseLengthStr failed, res != exp, %v != %v",
			res, exp,
		)
	}
}
