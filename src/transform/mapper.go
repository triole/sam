package transform

import (
	"reflect"
	"sort"
	"strings"

	toml "github.com/pelletier/go-toml"
)

type tFuncList []tFunc
type tFuncMap map[string]tFunc

type tFunc struct {
	Category string
	Command  string
	Args     string
	Desc     string
	Usage    interface{}
	Func     interface{}
}

func (fl tFuncList) Len() int {
	return len(fl)
}

func (fl tFuncList) Less(i, j int) bool {
	iSort := fl[i].Category + fl[i].Command
	jSort := fl[j].Category + fl[j].Command
	return iSort < jSort
}

func (fl tFuncList) Swap(i, j int) {
	fl[i], fl[j] = fl[j], fl[i]
}

func (tr *Transform) makeFuncList() {
	for key, val := range tr.FuncMap {
		val.Command = key
		tr.FuncList = append(tr.FuncList, val)
	}
	sort.Sort(tr.FuncList)
}

func (tr *Transform) makeFuncMap() {
	err := toml.Unmarshal(embedMapper, &tr.FuncMap)
	for key, val := range tr.FuncMap {
		val.Command = key
		tr.FuncMap[key] = val
	}
	logFatal(err, "error, can not unmarshal mapper toml")
}

func (tr Transform) Call() (result interface{}, err error) {
	if fn, ok := tr.FuncMap[tr.CLI.Command]; ok {
		fnct := reflect.ValueOf(&tr)
		meth := fnct.MethodByName(fn.Func.(string))
		methArgsInNo := meth.Type().NumIn()

		methArgsInArr := []string{strings.Join(tr.CLI.Args, " ")}
		if methArgsInNo > 1 {
			methArgsInArr = []string{
				tr.CLI.Args[0], strings.Join(tr.CLI.Args[1:], " "),
			}
		}

		methArgsInRef := make([]reflect.Value, methArgsInNo)
		for idx, param := range methArgsInArr {
			methArgsInRef[idx] = reflect.ValueOf(param)
		}

		var res []reflect.Value = meth.Call(methArgsInRef)
		result = res[0].Interface()
	}
	return
}
