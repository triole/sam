package transform

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// ListFunctions prints all available string transformation functions
func (tr Transform) ListFunctions() {
	t := table.NewWriter()
	t.SetStyle(table.Style{
		Name: "myNewStyle",
		Box: table.BoxStyle{
			BottomLeft:       "\\",
			BottomRight:      "/",
			BottomSeparator:  "v",
			Left:             "[",
			LeftSeparator:    "{",
			MiddleHorizontal: "-",
			MiddleSeparator:  "+",
			MiddleVertical:   "|",
			PaddingLeft:      " ",
			PaddingRight:     " ",
			Right:            " ]",
			RightSeparator:   "}",
			TopLeft:          "(",
			TopRight:         ")",
			TopSeparator:     "^",
			UnfinishedRow:    " ~~~",
		},
		Format: table.FormatOptions{
			Header: text.FormatUpper,
		},
		Options: table.Options{
			DrawBorder:      false,
			SeparateColumns: true,
			SeparateFooter:  true,
			SeparateHeader:  true,
			SeparateRows:    false,
		},
	})

	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1},
		{Number: 2},
		{Number: 3},
	})

	t.SetOutputMirror(os.Stdout)
	if tr.CLI.ListShort {
		t.AppendHeader(table.Row{
			"category", "command", "args", "usage",
		})
		for _, el := range tr.FuncList {
			t.AppendRow(
				[]interface{}{
					el.Category, el.Command, el.Args,
					printUsage(el),
				},
			)
		}
	} else {
		t.AppendHeader(table.Row{
			"category", "command", "args", "description", "usage",
		})
		for _, el := range tr.FuncList {
			t.AppendRow(
				[]interface{}{
					el.Category, el.Command, el.Args, el.Desc,
					printUsage(el),
				},
			)
		}
	}

	fmt.Printf("\n")
	t.Render()
	fmt.Printf("\n")
}

func printUsage(fnct tFunc) (r string) {
	var arr []string
	switch val := fnct.Usage.(type) {
	case string:
		arr = append(arr, val)
	case []interface{}:
		for _, el := range val {
			arr = append(arr, el.(string))
		}
	}
	for idx, el := range arr {
		suf := "\n"
		if idx+1 == len(arr) {
			suf = ""
		}
		r += "sam " + fnct.Command + " " + el + suf
	}
	return r
}
