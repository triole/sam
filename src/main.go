package main

import (
	"fmt"
	"sam/src/transform"
	"strings"
)

func main() {
	parseArgs()
	tr := transform.Init()

	println(CLI.Command)
	fmt.Printf("%q\n", CLI.StringToTransform)
	stringToTransform := strings.Join(CLI.StringToTransform, " ")
	fmt.Printf("%q\n", stringToTransform)
	if stringToTransform == "" {
		stringToTransform = getStdin()
	}

	fmt.Printf("%s\n", tr.Title(stringToTransform))
}
