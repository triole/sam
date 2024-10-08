package transform

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func separateFirstArg(s string) (string, string) {
	arr := strings.Split(s, " ")
	return arr[0], strings.Join(arr[1:], " ")
}

func parseLengthStr(s string) (i int) {
	arr := rxSplitToFloat("[a-zA-Z]+", s)
	if len(arr) > 1 && strings.Contains(s, "e") {
		i = int(arr[0] * math.Pow(10, arr[1]))
	}
	if len(arr) > 1 && strings.Contains(s, "p") {
		i = int(math.Pow(arr[0], arr[1]))
	}
	if len(arr) <= 1 {
		i = int(arr[0])
	}
	return
}

func rxSplitToFloat(rx, txt string) (arr []float64) {
	re := regexp.MustCompile(rx)
	split := re.Split(txt, -1)
	for i := range split {
		fl, err := strconv.ParseFloat(split[i], 32)
		if err == nil {
			arr = append(arr, fl)
		}
	}
	return
}

func logFatal(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}
