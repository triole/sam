package transform

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

// Transform holds the class
type Transform struct{}

func (tr Transform) makeFuncMap() (funcMap map[string]interface{}) {
	funcMap = make(map[string]interface{})
	funcMap["title"] = tr.title
	funcMap["md5"] = tr.md5
	funcMap["sha1"] = tr.sha1
	funcMap["sha256"] = tr.sha256
	funcMap["sha512"] = tr.sha512
	return
}

func (tr Transform) getFunc(funcName string) (r interface{}) {
	funcMap := tr.makeFuncMap()
	if val, ok := funcMap[funcName]; ok {
		r = val
	}
	return
}

// Init does what it says, it initialises the transform class
func Init() (tr Transform) {
	return Transform{}
}

// Call calls all available functions
func (tr Transform) Call(funcName string, params ...interface{}) (result interface{}, err error) {
	fn := tr.getFunc(funcName)
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
func (tr Transform) ListFunctions() {
	funcMap := tr.makeFuncMap()
	var iterator []string
	for key := range funcMap {
		iterator = append(iterator, key)
	}
	sort.Strings(iterator)
	fmt.Printf("\n%s\n", "String transformation command not found.")
	fmt.Printf("%s\n", "Please use one of the available:")
	for _, el := range iterator {
		fmt.Printf("    %s\n", el)
	}
}
