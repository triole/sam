package transform

import (
	"sam/src/conf"
	"sam/src/implant"
	"testing"
)

func TestColor(t *testing.T) {
	validateTestColor("#222222", []uint8{34, 34, 34}, t)
	validateTestColor("#999", []uint8{153, 153, 153}, t)
	validateTestColor("999", []uint8{153, 153, 153}, t)
	validateTestColor("#336699", []uint8{51, 102, 153}, t)
	validateTestColor("336699", []uint8{51, 102, 153}, t)
	validateTestColor("#159", []uint8{17, 85, 153}, t)
	validateTestColor("34 34 34", []uint8{34, 34, 34}, t)
}

func validateTestColor(inp string, exp []uint8, t *testing.T) {
	tr := Init(conf.New(), implant.Init())
	cc := tr.colorInfo(inp)
	if cc.RGB[0] != exp[0] || cc.RGB[1] != exp[1] || cc.RGB[2] != exp[2] {
		t.Errorf(
			"color info test failed: %+v, got: %+v",
			cc.RGB, exp,
		)
	}
}
