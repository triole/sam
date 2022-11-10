package transform

import (
	"errors"
	"reflect"
	"sort"

	"github.com/pelletier/go-toml"
)

type tFuncList []tFunc
type tFuncMap map[string]tFunc

type tFunc struct {
	Category string
	Command  string
	Args     string
	Desc     string
	Usage    string
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
	logFatal(err, "Error unmarshal mapper toml")
}

func (tr Transform) Call(funcName string, params ...interface{}) (result interface{}, err error) {
	if fn, ok := tr.FuncMap[funcName]; ok {
		fnct := reflect.ValueOf(&tr)
		meth := fnct.MethodByName(fn.Func.(string))

		if len(params) != meth.Type().NumIn() {
			err = errors.New("Number of params does not fit")
			return
		}

		in := make([]reflect.Value, len(params))
		for k, param := range params {
			in[k] = reflect.ValueOf(param)
		}

		var res []reflect.Value
		res = meth.Call(in)
		result = res[0].Interface()
	}
	return
}
