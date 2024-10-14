package transform

import (
	"fmt"
	"math"
	"strconv"

	calculator "github.com/mnogu/go-calculator"
)

func (tr Transform) runCalc() (r string) {
	val, err := calculator.Calculate(tr.Conf.String)
	logFatal(err, "calc failed: ")
	if tr.Conf.Precision >= 0 {
		val = round(val, uint(tr.Conf.Precision))
	}
	form := "%f"
	if tr.Conf.Precision >= 0 {
		form = "%0." + strconv.Itoa(tr.Conf.Precision) + "f"
	}
	r = fmt.Sprintf(form, val)
	return
}

func round(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
