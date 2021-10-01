package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"sam/src/transform"
	"sort"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

var (
	tableDescMaxWidth = 52
)

type tFuncList []tFunc
type tFuncMap map[string]tFunc

type tFunc struct {
	Name     string
	Desc     string
	Category string
	Func     interface{}
	Sorter   int
}

func (fl tFuncList) Len() int {
	return len(fl)
}

func (fl tFuncList) Less(i, j int) bool {
	if fl[i].Sorter == fl[j].Sorter {
		return fl[i].Name < fl[j].Name
	}
	return fl[i].Sorter < fl[j].Sorter
}

func (fl tFuncList) Swap(i, j int) {
	fl[i], fl[j] = fl[j], fl[i]
}

func makeFuncMap() (fm tFuncMap) {
	tr := transform.Init()
	fm = make(tFuncMap)
	fm = addToMap(fm, tr.Title, "title", "title case", "case", 0)
	fm = addToMap(fm, tr.LowerCase, "lower", "to lowercase", "case", 0)
	fm = addToMap(fm, tr.UpperCase, "upper", "to uppercase", "case", 0)
	fm = addToMap(fm, tr.SnakeCase, "snake", "to snakecase", "case", 0)
	fm = addToMap(fm, tr.CamelCase, "camel", "to camelcase", "case", 0)

	fm = addToMap(
		fm, tr.Bool, "bool",
		"return boolean: 1, enable, enabled, on and true return true, "+
			"everything else false (case doesn't matter)",
		"logical", 1,
	)

	fm = addToMap(fm, tr.FromBase64, "fr_b64", "from base64 to string", "encoding", 2)
	fm = addToMap(fm, tr.ToBase64, "to_b64", "to base64 from string", "encoding", 2)

	fm = addToMap(fm, tr.Md5, "md5", "md5 hash", "hash", 3)
	fm = addToMap(fm, tr.Sha1, "sha1", "sha1 hash", "hash", 3)
	fm = addToMap(fm, tr.Sha256, "sha256", "sha256 hash", "hash", 3)
	fm = addToMap(fm, tr.Sha512, "sha512", "sha512 hash", "hash", 3)

	fm = addToMap(
		fm, tr.DirName, "folder",
		"folder of a path string, return everything up to last path separator, "+
			"path separators trailing the input are ignored "+
			"(i.e. /tmp/hello/ -> /tmp)",
		"path", 4,
	)
	fm = addToMap(
		fm, tr.TidyFileName1, "tp1",
		"tidy path 1, remove multiple path separators",
		"path", 4,
	)
	fm = addToMap(
		fm, tr.TidyFileName2, "tp2",
		"as tp1, but also remove all accents, then replace characters not being "+
			"alpha numerics, dashes, underscores or path separators by underscores",
		"path", 4,
	)
	fm = addToMap(
		fm, tr.TidyFileName3, "tp3",
		"as tp2, plus lower case conversion",
		"path", 4,
	)
	fm = addToMap(
		fm, tr.TidyFileName4, "tp4",
		"as tp3, replace double underscores which may appear during conversion by a single one",
		"path", 4,
	)
	return
}

func addToMap(fm tFuncMap, f interface{}, name, desc, category string, sorter int) tFuncMap {
	fm[name] = newFunc(f, name, desc, category, sorter)
	return fm
}

func newFunc(function interface{}, name, desc, category string, sorter int) tFunc {
	return tFunc{
		Name:     name,
		Desc:     desc,
		Category: category,
		Func:     function,
		Sorter:   sorter,
	}
}

// Call calls all available functions
func Call(funcName string, params ...interface{}) (result interface{}, err error) {
	fn := getFunc(makeFuncMap(), funcName)
	if fn != nil {
		f := reflect.ValueOf(fn)
		if len(params) != f.Type().NumIn() {
			err = errors.New("The number of params is out of index")
			return
		}
		in := make([]reflect.Value, len(params))
		for k, param := range params {
			in[k] = reflect.ValueOf(param)
		}
		var res []reflect.Value
		res = f.Call(in)
		result = res[0].Interface()
	}
	return
}

// ListFunctions prints all available string transformation functions
func ListFunctions() {
	fm := makeFuncMap()
	var fl tFuncList
	for _, val := range fm {
		fl = append(fl, val)
	}
	sort.Sort(tFuncList(fl))
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
			SeparateRows:    true,
		},
	})

	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1},
		{Number: 2},
		{Number: 3, WidthMax: tableDescMaxWidth},
	})

	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"command", "category", "description",
	})
	for _, el := range fl {
		t.AppendRow(
			[]interface{}{
				el.Name, fm[el.Name].Category, wordWrap(fm[el.Name].Desc, tableDescMaxWidth),
			},
		)
	}
	fmt.Printf("\n")
	t.Render()
	fmt.Printf("\n")
}

func getFunc(fm tFuncMap, funcName string) (r interface{}) {
	if val, ok := fm[funcName]; ok {
		r = val.Func
	}
	return
}
