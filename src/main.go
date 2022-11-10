package main

import (
	"fmt"
	"log"
	"sam/src/transform"
	"strings"
)

func main() {
	parseArgs()

	args := strings.Join(CLI.Args, " ")
	if args == "" {
		args = getStdin()
	}

	tr := transform.Init(CLI.List, CLI.ListShort)

	res, err := tr.Call(CLI.Command, args)
	if err != nil {
		log.Fatalf("Error calling command: %+v\n", err.Error())
	}
	if res != nil {
		fmt.Printf("%s\n", res)
	} else {
		tr.ListFunctions()
		if CLI.List == false && CLI.ListShort == false {
			fmt.Printf("%s\n\n",
				"String transformation command not found. "+
					"Please use one of the above.",
			)
		}
	}
}
