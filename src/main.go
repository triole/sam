package main

import (
	"fmt"
	"sam/src/transform"
	"strings"
)

func main() {
	parseArgs()
	tr := transform.Init()

	stringToTransform := strings.Join(CLI.StringToTransform, " ")
	if stringToTransform == "" {
		stringToTransform = getStdin()
	}

	res, _ := tr.Call(CLI.Command, stringToTransform)
	if res != nil {
		fmt.Printf("%s\n", res)
	} else {
		tr.ListFunctions()
	}
}
