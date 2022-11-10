package transform

import (
	_ "embed"
)

//go:embed mapper.toml
var embedMapper []byte

// Transform holds the class
type Transform struct {
	FuncList     tFuncList
	FuncMap      tFuncMap
	CLIList      bool
	CLIListShort bool
}

// Init does what it says, it initialises the transform class
func Init(CLIList bool, CLIListShort bool) (tr Transform) {
	tr = Transform{
		FuncList:     tFuncList{},
		FuncMap:      make(tFuncMap),
		CLIList:      CLIList,
		CLIListShort: CLIListShort,
	}
	tr.makeFuncMap()
	tr.makeFuncList()
	return
}
