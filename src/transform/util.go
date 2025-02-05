package transform

import (
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

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

func rxCompile(str string) (r *regexp.Regexp) {
	var err error
	r, err = regexp.Compile(str)
	logFatal(err, "can not compile regex")
	return
}

// func rxFind(rx string, content string) (r string) {
// 	temp := rxCompile(rx)
// 	r = temp.FindString(content)
// 	return
// }

func rxMatch(rx string, str string) (b bool) {
	re := rxCompile(rx)
	b = re.MatchString(str)
	return
}

func rxSub(str string, rx string, rep string) (r string) {
	re := regexp.MustCompile(rx)
	r = re.ReplaceAllString(str, rep)
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

func stringToFloat(str string) (fl float64, err error) {
	return strconv.ParseFloat(str, 32)
}

func printTable(arr [][]interface{}) {
	t := table.NewWriter()
	t.SetStyle(table.Style{
		Box: table.BoxStyle{
			MiddleVertical: "|",
			PaddingLeft:    " ",
			PaddingRight:   " ",
		},
		Options: table.Options{
			DrawBorder:      false,
			SeparateColumns: false,
			SeparateFooter:  false,
			SeparateHeader:  false,
			SeparateRows:    false,
		},
		Format: table.FormatOptions{
			Footer: text.FormatUpper,
			Header: text.FormatUpper,
			Row:    text.FormatDefault,
		},
	})
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(arr[0])
	for _, el := range arr[1:] {
		t.AppendRow(el)
	}
	println()
	t.Render()
}

func logFatal(err error, msg string) {
	if err != nil {
		log.Fatalf("[fatal] %s: %s", msg, err.Error())
	}
}
