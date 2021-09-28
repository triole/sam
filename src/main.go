package main

import (
	"fmt"
	"strings"
)

func main() {
	parseArgs()

	stringToTransform := strings.Join(CLI.StringToTransform, " ")
	if stringToTransform == "" {
		stringToTransform = getStdin()
	}

	res, _ := Call(CLI.Command, stringToTransform)
	if res != nil {
		fmt.Printf("%s\n", res)
	} else {
		ListFunctions()
	}
}
