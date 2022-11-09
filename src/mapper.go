package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"sam/src/transform"
	"sort"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// var (
// 	tableDescMaxWidth = 52
// )

type tFuncList []tFunc
type tFuncMap map[string]tFunc

type tFunc struct {
	Category string
	Name     string
	Args     string
	Desc     string
	Usage    string
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
	fm = addToMap(
		fm, tr.TrimPrefix, "trim", "trimprefix", "prefix, str",
		"remove prefix, requires two args: string, prefix to remove",
		"_ _hello", 0,
	)
	fm = addToMap(
		fm, tr.TrimPrefixAggressive, "trim", "trimprefixag", "prefix, str",
		"trim prefix aggressive, remove multiple occurences of prefix",
		"_ __hello", 0,
	)
	fm = addToMap(
		fm, tr.TrimSuffix, "trim", "trimsuffix", "suffix, str",
		"like trimprefix but removing end of a string",
		"_ hello_", 0,
	)
	fm = addToMap(
		fm, tr.TrimSuffixAggressive, "trim", "trimsuffixag", "suffix, str",
		"like trim prefix aggressive but at the end",
		"_ hello__", 0,
	)
	fm = addToMap(
		fm, tr.TrimSpace, "trim", "trimspace", "str",
		"remove spaces or tabs around a string",
		"", 1,
	)

	fm = addToMap(
		fm, tr.RemoveMultiSpace, "replace", "rms", "str",
		"replace multiple spaces or tabs after one another by one space",
		"", 2,
	)

	fm = addToMap(fm, tr.Title, "case", "title", "str", "title case", "", 3)
	fm = addToMap(fm, tr.LowerCase, "case", "lower", "str", "to lowercase", "", 3)
	fm = addToMap(fm, tr.UpperCase, "case", "upper", "str", "to uppercase", "", 3)
	fm = addToMap(fm, tr.SnakeCase, "case", "snake", "str", "to snakecase", "", 3)
	fm = addToMap(fm, tr.CamelCase, "case", "camel", "str", "to camelcase", "", 3)

	fm = addToMap(
		fm, tr.Bool, "bool", "str", "logical",
		"return bool; true on: 1, enable, enabled, on, true; "+
			"false on everything else; case insensitive",
		"enabled", 3,
	)

	fm = addToMap(fm, tr.FromBase64, "encoding", "txt-b64", "str", "from base64 to string", "", 4)
	fm = addToMap(fm, tr.ToBase64, "encoding", "b64-txt", "str", "to base64 from string", "", 4)

	fm = addToMap(fm, tr.FromURL, "encoding", "url-txt", "str", "from url to plain string", "", 4)
	fm = addToMap(fm, tr.ToURL, "encoding", "txt-url", "str", "to url from plain string", "", 4)

	fm = addToMap(fm, tr.Md5, "hash", "md5", "str", "md5 hash", "", 5)
	fm = addToMap(fm, tr.Sha1, "hash", "sha1", "str", "sha1 hash", "", 5)
	fm = addToMap(fm, tr.Sha256, "hash", "sha256", "str", "sha256 hash", "", 5)
	fm = addToMap(fm, tr.Sha512, "hash", "sha512", "str", "sha512 hash", "", 5)
	fm = addToMap(
		fm, tr.Blake3, "hash", "blake3",
		"size, str", "blake3 hash, flexible hash size", "128 hello", 5,
	)
	fm = addToMap(fm, tr.Ripemd160, "hash", "ripemd160", "str", "ripemd160 hash", "", 5)
	fm = addToMap(fm, tr.Whirlpool, "hash", "whirlpool", "str", "whirlpool hash", "", 5)

	fm = addToMap(
		fm, tr.DirName, "path", "folder", "str",
		"folder of a path, everything up to last path separator, trailing path separators ignored",
		"", 6,
	)
	fm = addToMap(
		fm, tr.TidyFilePath, "path", "tfn", "str",
		"tidy file name, only allow '[0-9a-z\\-_]', replace multiple path separators, underscores and dashes",
		"", 6,
	)
	fm = addToMap(
		fm, tr.TidyPathSeparators, "path", "tps", "str",
		"tidy path separators, replace multiple path separators after one another",
		"", 6,
	)
	return
}

func addToMap(fm tFuncMap, f interface{}, category, name, args, desc, usage string, sorter int) tFuncMap {
	fm[name] = newFunc(f, category, name, args, desc, usage, sorter)
	return fm
}

func newFunc(function interface{}, category, name, args, desc, usage string, sorter int) tFunc {
	return tFunc{
		Category: category,
		Name:     name,
		Desc:     desc,
		Args:     args,
		Usage:    usage,
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
			err = errors.New("Number of params is out of index")
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
			SeparateRows:    false,
		},
	})

	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1},
		{Number: 2},
		{Number: 3},
	})

	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"category", "command", "args", "description", "usage",
	})
	for _, el := range fl {
		t.AppendRow(
			[]interface{}{
				el.Category, el.Name, el.Args, el.Desc,
				printUsage(el),
			},
		)
	}
	fmt.Printf("\n")
	t.Render()
	fmt.Printf("\n")
}

func printUsage(el tFunc) (r string) {
	if el.Usage != "" {
		r = "sam " + el.Name + " " + el.Usage
	}
	return r
}

func getFunc(fm tFuncMap, funcName string) (r interface{}) {
	if val, ok := fm[funcName]; ok {
		r = val.Func
	}
	return
}
