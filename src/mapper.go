package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"sam/src/transform"
	"sort"

	"github.com/jedib0t/go-pretty/table"
)

type tFuncList []tFunc
type tFuncMap map[string]tFunc

type tFunc struct {
	Name   string
	Desc   string
	Func   interface{}
	Sorter int
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
	fm = addToMap(fm, tr.Title, "title", "title case", 0)
	fm = addToMap(fm, tr.Lowercase, "lowercase", "to lowercase", 0)
	fm = addToMap(fm, tr.Uppercase, "uppercase", "to uppercase", 0)
	fm = addToMap(fm, tr.Md5, "md5", "md5 hash", 1)
	fm = addToMap(fm, tr.Sha1, "sha1", "sha1 hash", 1)
	fm = addToMap(fm, tr.Sha256, "sha256", "sha256 hash", 1)
	fm = addToMap(fm, tr.Sha512, "sha512", "sha512 hash", 1)
	return
}

func addToMap(fm tFuncMap, f interface{}, name, desc string, sorter int) tFuncMap {
	fm[name] = newFunc(f, name, desc, sorter)
	return fm
}

func newFunc(function interface{}, name, desc string, sorter int) tFunc {
	return tFunc{
		Name:   name,
		Desc:   desc,
		Func:   function,
		Sorter: sorter,
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
	fmt.Printf("\n%s\n", "String transformation command not found.")
	fmt.Printf("%s\n", "Please use one of the available...")
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
		Options: table.Options{
			DrawBorder:      false,
			SeparateColumns: true,
			SeparateFooter:  true,
			SeparateHeader:  true,
			SeparateRows:    false,
		},
	})
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"name", "description",
	})
	var lastSorter int
	var currentSorter int
	for _, el := range fl {
		lastSorter = currentSorter
		currentSorter = el.Sorter
		if lastSorter != currentSorter {
			t.AppendRow([]interface{}{""})
		}
		t.AppendRow(
			[]interface{}{
				el.Name, fm[el.Name].Desc,
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
