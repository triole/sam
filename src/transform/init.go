package transform

import (
	_ "embed"
)

//go:embed mapper.toml
var embedMapper []byte

// Transform holds the class
type Transform struct {
	FuncList tFuncList
	FuncMap  tFuncMap
	CLI      TransformCLI
}

type TransformCLI struct {
	Command   string
	Args      []string
	File      string
	List      bool
	ListShort bool
}

// Init does what it says, it initialises the transform class
func Init(tcli TransformCLI) (tr Transform) {
	tr = Transform{
		FuncList: tFuncList{},
		FuncMap:  make(tFuncMap),
		CLI:      tcli,
	}
	tr.makeFuncMap()
	tr.makeFuncList()
	return
}
