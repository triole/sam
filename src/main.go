package main

import (
	"fmt"
	"strings"
)

func main() {
	parseArgs()

	args := strings.Join(CLI.Args, " ")
	if args == "" {
		args = getStdin()
	}

	res, _ := Call(CLI.Command, args)
	if res != nil {
		fmt.Printf("%s\n", res)
	} else {
		ListFunctions()
		if CLI.List == false {
			fmt.Printf("%s\n\n",
				"String transformation command not found. "+
					"Please use one of the above.",
			)
		}
	}
}
