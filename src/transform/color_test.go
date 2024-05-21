package transform

import (
	"testing"
)

func TestColorInfo(t *testing.T) {
	validateTestColorInfo("#222222", []uint8{34, 34, 34}, t)
	validateTestColorInfo("#999", []uint8{153, 153, 153}, t)
	validateTestColorInfo("999", []uint8{153, 153, 153}, t)
	validateTestColorInfo("#336699", []uint8{51, 102, 153}, t)
	validateTestColorInfo("336699", []uint8{51, 102, 153}, t)
	validateTestColorInfo("#159", []uint8{17, 85, 153}, t)
	validateTestColorInfo("34 34 34", []uint8{34, 34, 34}, t)
}

func validateTestColorInfo(inp string, exp []uint8, t *testing.T) {
	cc := tr.colorInfo(inp)
	if cc.RGB[0] != exp[0] || cc.RGB[1] != exp[1] || cc.RGB[2] != exp[2] {
		t.Errorf(
			"color info test failed: %+v, got: %+v",
			cc.RGB, exp,
		)
	}
}
